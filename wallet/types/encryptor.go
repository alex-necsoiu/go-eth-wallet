package types

// Encryptor is the interface for encrypting and decrypting sensitive information in wallets.
type Encryptor interface {
	// Name() provides the name of the encryptor.
	Name() string

	// Version() provides the version of the encryptor.
	Version() uint

	// String provides a string value for the encryptor.
	String() string

	// Encrypt encrypts a byte array with its encryption mechanism and key.
	Encrypt(data []byte, key string) (map[string]any, error)

	// Decrypt encrypts a byte array with its encryption mechanism and key.
	Decrypt(data map[string]any, key string) ([]byte, error)
}
