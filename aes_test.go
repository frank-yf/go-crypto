package crypto_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/frank-yf/go-crypto"
)

var (
	origData      = "abcd"
	origDataBytes = []byte(origData)
	key           = "0123456789abcdef"
	keyBytes      = []byte(key)

	dataSize = 1000
)

type bytesStruct struct {
	data   []byte
	crypto *crypto.AesECB
}

type stringStruct struct {
	data   string
	crypto *crypto.AesECB
}

func cryptoBytes(t TestingAny) func([]byte, *crypto.AesECB) {
	return func(bs []byte, c *crypto.AesECB) {
		encrypt := c.Encrypt(bs)
		decrypt := c.Decrypt(encrypt)
		equalLength(t, bs, decrypt)
	}
}

func TestAesECBCrypto(t *testing.T) {
	ts := bytesTable(t)
	for _, tt := range ts {
		cryptoBytes(t)(tt.data, tt.crypto)
	}
}

func TestAesECBCrypto_WithPool(t *testing.T) {
	ts := bytesTable(t, crypto.WithPool(1024))
	for _, tt := range ts {
		cryptoBytes(t)(tt.data, tt.crypto)
	}
}

func cryptoString(t TestingAny) func(string, *crypto.AesECB) {
	return func(s string, c *crypto.AesECB) {
		toHex := c.EncryptToHex(s)

		decrypted, err := c.DecryptFromHex(toHex)
		noError(t, err)
		equalString(t, s, decrypted)

		toBase64 := c.EncryptToBase64(s)

		decrypted, err = c.DecryptFromBase64(toBase64)
		noError(t, err)
		equalString(t, s, decrypted)
	}
}

func TestAesECBCryptoString(t *testing.T) {
	ts := stringTable(t)
	for _, tt := range ts {
		cryptoString(t)(tt.data, tt.crypto)
	}
}

func TestAesECBCryptoString_WithPool(t *testing.T) {
	ts := stringTable(t, crypto.WithPool(1024))
	for _, tt := range ts {
		cryptoString(t)(tt.data, tt.crypto)
	}
}

func bytesTable(t TestingAny, opts ...crypto.AesECBOptions) []bytesStruct {
	cry, err := crypto.NewAesECB(key)
	noError(t, err)

	for _, opt := range opts {
		opt(cry)
	}

	tests := make([]bytesStruct, 0, dataSize)
	for i := 0; i < dataSize; i++ {
		tests = append(tests, bytesStruct{
			data:   bytes.Repeat(origDataBytes, i),
			crypto: cry,
		})
	}

	return tests
}

func stringTable(t TestingAny, opts ...crypto.AesECBOptions) []stringStruct {
	cry, err := crypto.NewAesECB(key)
	noError(t, err)

	for _, opt := range opts {
		opt(cry)
	}

	tests := make([]stringStruct, 0, dataSize)
	for i := 0; i < dataSize; i++ {
		tests = append(tests, stringStruct{
			data:   strings.Repeat(origData, i),
			crypto: cry,
		})
	}

	return tests
}
