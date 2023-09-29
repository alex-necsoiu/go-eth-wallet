package unencrypted

// Encryptor is an encryptor that stores data unencrypted.
type Encryptor struct{}

type unencrypted struct {
	Key string `json:"key"`
}

// New creates a new null encryptor.
func New() *Encryptor {
	return &Encryptor{}
}

// String returns the string for this encryptor.
func (e *Encryptor) String() string {
	return "unencryptedv1"
}

// Name returns the name of this encryptor.
func (e *Encryptor) Name() string {
	return "unencrypted"
}

// Version returns the version of this encryptor.
func (e *Encryptor) Version() uint {
	return 1
}
