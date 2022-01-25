// Copy for github.com/gin-gonic/gin/internal/bytesconv
package crypto

import (
	"encoding/base64"
	"encoding/hex"
	"unsafe"
)

// StringToBytes converts string to byte slice without a memory allocation.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func noCopyHexEncode(src []byte) string {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return BytesToString(dst)
}

func noCopyBase64Encode(src []byte) string {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(buf, src)
	return BytesToString(buf)
}

func noCopyBase64Decode(s string) ([]byte, error) {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(s)))
	n, err := base64.StdEncoding.Decode(dst, StringToBytes(s))
	return dst[:n], err
}
