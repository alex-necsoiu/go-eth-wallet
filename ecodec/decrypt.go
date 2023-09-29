package ecodec

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"errors"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

// Decrypt decrypts some data given a secret.
func Decrypt(encryptedData []byte, key []byte) ([]byte, error) {
	if len(encryptedData) < _versionLen+_saltLen+_ivLen+_checksumLen {
		return nil, fmt.Errorf("encrypted data must be at least %d bytes", _versionLen+_saltLen+_ivLen+_checksumLen)
	}

	if encryptedData[0] != version {
		return nil, fmt.Errorf("unhandled version %#02x", encryptedData[0])
	}

	// Unpack data
	salt := encryptedData[_versionLen : _versionLen+_saltLen]
	iv := encryptedData[_versionLen+_saltLen : _versionLen+_saltLen+_ivLen]
	checksum := encryptedData[_versionLen+_saltLen+_ivLen : _versionLen+_saltLen+_ivLen+_checksumLen]
	data := encryptedData[_versionLen+_saltLen+_ivLen+_checksumLen:]

	decryptionKey := pbkdf2.Key(key, salt, _pbkdf2c, _pbkdf2KeyLen, sha256.New)

	h := sha256.New()
	_, err := h.Write(decryptionKey[16:32])
	if err != nil {
		return nil, err
	}
	_, err = h.Write(encryptedData[_versionLen+_saltLen+_ivLen+_checksumLen:])
	if err != nil {
		return nil, err
	}
	calculatedChecksum := h.Sum(nil)

	if !bytes.Equal(calculatedChecksum, checksum) {
		return nil, errors.New("invalid key")
	}

	res := make([]byte, len(data))
	aesCipher, err := aes.NewCipher(decryptionKey[:16])
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(aesCipher, iv)
	stream.XORKeyStream(res, data)

	return res, nil
}
