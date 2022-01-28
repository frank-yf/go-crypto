package codec

import (
"encoding/base64"
)

// Base64EncodeToString encode base64 to string
// minimal use of copy
func Base64EncodeToString(src []byte) string {
	return BytesToString(Base64Encode(src))
}

// Base64DecodeFromString decode base64 of string
// minimal use of copy
func Base64DecodeFromString(s string) ([]byte, error) {
	return Base64Decode(StringToBytes(s))
}

// Base64Decode decode base64 of []byte
func Base64Decode(bs []byte) ([]byte, error) {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(bs)))
	n, err := base64.StdEncoding.Decode(dst, bs)
	return dst[:n], err
}

// Base64Encode decode base64 of []byte
func Base64Encode(src []byte) []byte {
	encode := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(encode, src)
	return encode
}
