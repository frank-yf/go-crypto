package codec

import (
	"encoding/hex"
)

// HexEncodeToString encode hex to string
func HexEncodeToString(src []byte) string {
	return BytesToString(HexEncode(src))
}

// HexDecodeFromString decode hex from string
func HexDecodeFromString(s string) (bs []byte, err error) {
	return hex.DecodeString(s)
}

// HexDecode decode hex
func HexDecode(bs []byte) (dst []byte, err error) {
	dst = make([]byte, hex.DecodedLen(len(bs)))
	_, err = hex.Decode(dst, bs)
	return
}

// HexEncode encode hex
func HexEncode(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}
