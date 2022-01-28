package crypto

type StringEncoder func([]byte) string

type StringDecoder func(string) ([]byte, error)

type Crypto interface {
	Encrypt([]byte) []byte

	EncryptTo(string, StringEncoder) string

	Decrypt([]byte) []byte

	DecryptFrom(string, StringDecoder) (string, error)
}
