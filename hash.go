package crypto

import (
	"crypto/md5"
	"hash"
	"sync"

	"github.com/frank-yf/go-crypto/codec"
)

var (
	md5Hasher = NewHash(func() hash.Hash {
		return md5.New()
	})
)

func MD5(data []byte) []byte {
	h := md5.New()
	h.Write(data)
	return h.Sum(nil)
	// return md5Hasher.Sum(data) // 性能弱于基础包
}

func MD5Hex(data string) []byte {
	sum := MD5(codec.StringToBytes(data))
	return codec.HexEncode(sum)
}

func MD5HexToString(data string) string {
	sum := MD5(codec.StringToBytes(data))
	return codec.HexEncodeToString(sum)
}

type Hash struct {
	hashPool sync.Pool
}

func (h *Hash) getHash() (hash.Hash, func()) {
	digest := h.hashPool.Get().(hash.Hash)
	digest.Reset()

	return digest, func() {
		h.hashPool.Put(digest)
	}
}

func NewHash(f func() hash.Hash) *Hash {
	return &Hash{
		hashPool: sync.Pool{
			New: func() interface{} {
				return f()
			},
		},
	}
}

func (h *Hash) Sum(data []byte) []byte {
	hasher, done := h.getHash()
	defer done()

	hasher.Write(data)
	return hasher.Sum(nil)
}

func (h *Hash) SumString(data string, encoder StringEncoder) string {
	digest, done := h.getHash()
	defer done()

	digest.Write(codec.StringToBytes(data))
	return encoder(digest.Sum(nil))
}
