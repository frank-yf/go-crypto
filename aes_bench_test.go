package crypto_test

import (
	"bytes"
	"crypto/aes"
	"testing"

	"github.com/frank-yf/go-crypto"
)

func BenchmarkAesECBCrypto(b *testing.B) {
	f := func(data []byte, c *crypto.AesECB) {
		encrypt := c.Encrypt(data)
		decrypt := c.Decrypt(encrypt)
		equalLength(b, data, decrypt)
	}

	ts := bytesTable(b)
	for i := 0; i < b.N; i++ {
		for _, tt := range ts {
			f(tt.data, tt.crypto)
		}
	}
}

func BenchmarkAesECBCrypto_WithPool(b *testing.B) {
	f := func(data []byte, c *crypto.AesECB) {
		encrypt := c.Encrypt(data)
		decrypt := c.Decrypt(encrypt)
		equalLength(b, data, decrypt)
	}

	ts := bytesTable(b, crypto.WithPool(1024))
	for i := 0; i < b.N; i++ {
		for _, tt := range ts {
			f(tt.data, tt.crypto)
		}
	}
}

func BenchmarkAesECBCrypto_WithBase64(b *testing.B) {
	f := func(data string, c *crypto.AesECB) {
		toBase64 := c.EncryptToBase64(data)
		_, err := c.DecryptFromBase64(toBase64)
		noError(b, err)
	}

	ts := stringTable(b)
	for i := 0; i < b.N; i++ {
		for _, tt := range ts {
			f(tt.data, tt.crypto)
		}
	}
}

func BenchmarkAesECBCrypto_WithBase64_WithPool(b *testing.B) {
	f := func(data string, c *crypto.AesECB) {
		toBase64 := c.EncryptToBase64(data)
		_, err := c.DecryptFromBase64(toBase64)
		noError(b, err)
	}

	ts := stringTable(b, crypto.WithPool(1024))
	for i := 0; i < b.N; i++ {
		for _, tt := range ts {
			f(tt.data, tt.crypto)
		}
	}
}

func BenchmarkAesECBCrypto_WithHex(b *testing.B) {
	f := func(data string, c *crypto.AesECB) {
		toBase64 := c.EncryptToHex(data)
		_, err := c.DecryptFromHex(toBase64)
		noError(b, err)
	}

	ts := stringTable(b)
	for i := 0; i < b.N; i++ {
		for _, tt := range ts {
			f(tt.data, tt.crypto)
		}
	}
}

func BenchmarkAesECBCrypto_WithHex_WithPool(b *testing.B) {
	f := func(data string, c *crypto.AesECB) {
		toBase64 := c.EncryptToHex(data)
		_, err := c.DecryptFromHex(toBase64)
		noError(b, err)
	}

	ts := stringTable(b, crypto.WithPool(1024))
	for i := 0; i < b.N; i++ {
		for _, tt := range ts {
			f(tt.data, tt.crypto)
		}
	}
}

func BenchmarkNone(b *testing.B) {
	f := func(bs []byte, c *crypto.AesECB) {
		encrypt := EcbEncrypt(bs, keyBytes)
		EcbDecrypt(encrypt, keyBytes)
	}

	ts := bytesTable(b)
	for i := 0; i < b.N; i++ {
		for _, tt := range ts {
			f(tt.data, tt.crypto)
		}
	}
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func EcbDecrypt(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return PKCS7UnPadding(decrypted)
}

func EcbEncrypt(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	data = PKCS7Padding(data, block.BlockSize())
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}
