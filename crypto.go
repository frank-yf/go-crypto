package crypto

type EncryptEncoder func([]byte) string

type DecryptDecoder func(string) ([]byte, error)

type Crypto interface {
	Encrypt([]byte) []byte

	EncryptTo(string, EncryptEncoder) string

	Decrypt([]byte) []byte

	DecryptFrom(string, DecryptDecoder) (string, error)
}
