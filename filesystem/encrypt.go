package filesystem

import (
	"errors"

	"github.com/alex-necsoiu/go-eth-wallet/ecodec"
)

// encryptIfRequired encrypts data if required.
func (s *Store) encryptIfRequired(data []byte) ([]byte, error) {
	if len(data) == 0 {
		// No data means nothing to encrypt.
		return data, nil
	}

	if len(s.passphrase) == 0 {
		// No passphrase means nothing to encrypt with.
		return data, nil
	}

	if len(data) < 16 {
		return nil, errors.New("data must be at least 16 bytes")
	}

	var err error
	if data, err = ecodec.Encrypt(data, s.passphrase); err != nil {
		return nil, err
	}

	return data, nil
}

// decryptIfRequired decrypts data if required.
func (s *Store) decryptIfRequired(data []byte) ([]byte, error) {
	if len(data) == 0 {
		// No data means nothing to decrypt.
		return data, nil
	}

	if len(s.passphrase) == 0 {
		// No passphrase means nothing to decrypt with.
		return data, nil
	}

	if len(data) < 16 {
		return nil, errors.New("data must be at least 16 bytes")
	}

	var err error
	if data, err = ecodec.Decrypt(data, s.passphrase); err != nil {
		return nil, err
	}

	return data, nil
}
