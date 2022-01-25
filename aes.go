package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"sync"
)

var (
	emptyDone = func() {}
)

type AesECBOptions func(*AesECB)

type AesECB struct {
	ci cipher.Block

	pool *sync.Pool
}

func NewAesECB(key string, opts ...AesECBOptions) (c *AesECB, err error) {
	genKey := generateKey(StringToBytes(key))
	ci, err := aes.NewCipher(genKey)
	if err != nil {
		return
	}
	c = &AesECB{ci: ci}
	for _, opt := range opts {
		opt(c)
	}
	return
}

func WithPool(preSize int) AesECBOptions {
	return func(c *AesECB) {
		c.pool = &sync.Pool{
			New: func() interface{} {
				bs := make([]byte, preSize)
				return &bs
			},
		}
	}
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return
}

func (c *AesECB) makeBytes(length int, oneOff bool) ([]byte, func()) {
	if oneOff || c.pool == nil {
		return make([]byte, length), emptyDone
	}

	mb := *(c.pool.Get().(*[]byte)) // 池中取出复用数组不切分长度
	makeLen := length - len(mb)
	if makeLen > 0 {
		newBs := make([]byte, makeLen)
		mb = append(mb, newBs...)
	} else {
		mb = mb[:length] // 如果申请的数组长度没有全部覆盖，可能造成数据异常
	}
	return mb, func() { c.pool.Put(&mb) }
}

func (c *AesECB) Encrypt(origin []byte) (encrypted []byte) {
	blockSize := c.ci.BlockSize()
	plain := c.paddingPKCS7(origin, blockSize)

	encrypted, _ = c.makeBytes(len(plain), true) // 不能返回一个被复用的字节数组

	for bs, be := 0, blockSize; bs < len(plain); bs, be = bs+blockSize, be+blockSize {
		c.ci.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return
}

func (c *AesECB) EncryptTo(origin string, encoder EncryptEncoder) string {
	originBytes := StringToBytes(origin)

	blockSize := c.ci.BlockSize()
	plain := c.paddingPKCS7(originBytes, blockSize)

	encrypted, done := c.makeBytes(len(plain), false)
	defer done()

	for bs, be := 0, blockSize; bs < len(plain); bs, be = bs+blockSize, be+blockSize {
		c.ci.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return encoder(encrypted)
}

func (c *AesECB) EncryptToHex(origin string) string {
	return c.EncryptTo(origin, noCopyHexEncode)
}

func (c *AesECB) EncryptToBase64(origin string) string {
	return c.EncryptTo(origin, noCopyBase64Encode)
}

func (c *AesECB) Decrypt(encrypted []byte) []byte {
	decrypted, _ := c.makeBytes(len(encrypted), true) // 不能返回一个被复用的字节数组

	blockSize := c.ci.BlockSize()
	for bs, be := 0, blockSize; bs < len(encrypted); bs, be = bs+blockSize, be+blockSize {
		c.ci.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	return c.unPaddingPKCS7(decrypted)
}

func (c *AesECB) DecryptFrom(encrypted string, decoder DecryptDecoder) (decrypted string, err error) {
	encryptedBytes, err := decoder(encrypted)
	if err != nil {
		return
	}

	decryptedBytes, done := c.makeBytes(len(encryptedBytes), false)
	defer done()

	blockSize := c.ci.BlockSize()
	for bs, be := 0, blockSize; bs < len(encryptedBytes); bs, be = bs+blockSize, be+blockSize {
		c.ci.Decrypt(decryptedBytes[bs:be], encryptedBytes[bs:be])
	}

	copy(encryptedBytes, decryptedBytes)
	decrypted = BytesToString(c.unPaddingPKCS7(encryptedBytes))
	return
}

func (c *AesECB) paddingPKCS7(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	pad, done := c.makeBytes(padding, false)
	defer done()

	b := byte(padding)
	for i := 0; i < padding; i++ {
		pad[i] = b
	}
	return append(src, pad...)
}

func (c *AesECB) unPaddingPKCS7(pad []byte) []byte {
	length := len(pad)
	unPadding := int(pad[length-1])
	return pad[:(length - unPadding)]
}

func (c *AesECB) DecryptFromHex(encrypted string) (decrypted string, err error) {
	// 解码只新建了一个字节数组，没有优化空间
	from := hex.DecodeString
	return c.DecryptFrom(encrypted, from)
}

func (c *AesECB) DecryptFromBase64(encrypted string) (decrypted string, err error) {
	from := noCopyBase64Decode
	return c.DecryptFrom(encrypted, from)
}
