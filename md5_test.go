package crypto_test

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/frank-yf/go-crypto"
)

func TestMD5(t *testing.T) {
	bs := []byte("12345")

	hash := md5.New()
	hash.Write(bs)

	sum := hash.Sum(nil)

	equal(t, sum, crypto.MD5(bs))
	equal(t, sum, crypto.MD5(bs))
}

func TestMD5HexToString(t *testing.T) {
	s := "12345"
	bs := []byte(s)

	hash := md5.New()
	hash.Write(bs)

	sum := fmt.Sprintf("%x", hash.Sum(nil))

	equal(t, sum, crypto.MD5HexToString(s))
	equal(t, sum, crypto.MD5HexToString(s))
}

func BenchmarkMD5(b *testing.B) {
	data := []byte("12345")

	table := make([][]byte, dataSize)
	for i := 0; i < dataSize; i++ {
		table[i] = bytes.Repeat(data, i)
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range table {
			crypto.MD5(tt)
		}
	}
}

func BenchmarkMD5_None(b *testing.B) {
	data := []byte("12345")

	table := make([][]byte, dataSize)
	for i := 0; i < dataSize; i++ {
		table[i] = bytes.Repeat(data, i)
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range table {
			hash := md5.New()
			hash.Write(tt)
			hash.Sum(nil)
		}
	}
}

func BenchmarkMD5Hex(b *testing.B) {
	data := "12345"

	table := make([]string, dataSize)
	for i := 0; i < dataSize; i++ {
		table[i] = strings.Repeat(data, i)
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range table {
			crypto.MD5Hex(tt)
		}
	}
}

func BenchmarkMD5Hex_None(b *testing.B) {
	data := "12345"

	table := make([]string, dataSize)
	for i := 0; i < dataSize; i++ {
		table[i] = strings.Repeat(data, i)
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range table {
			hash := md5.New()
			_, _ = io.WriteString(hash, tt)
			_ = fmt.Sprintf("%x", hash.Sum(nil))
		}
	}
}
