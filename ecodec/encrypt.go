// Copyright © 2019 Weald Technology Trading
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ecodec

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"

	"golang.org/x/crypto/pbkdf2"
)

const (
	version = byte(1)

	_versionLen  = 1
	_saltLen     = 32
	_ivLen       = 16
	_checksumLen = 32

	_pbkdf2c      = 262144
	_pbkdf2KeyLen = 32
)

// Encrypt encrypts some data given a secret.
func Encrypt(data []byte, key []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("no data")
	}
	if len(key) == 0 {
		return nil, errors.New("no key")
	}
	if len(data) < 16 {
		return nil, errors.New("data must be at least 16 bytes")
	}

	encryptedData := make([]byte, len(data)+_versionLen+_saltLen+_ivLen+_checksumLen)
	encryptedData[0] = version

	// Random salt
	salt := make([]byte, _saltLen)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	copy(encryptedData[_versionLen:], salt)

	// Random IV
	iv := make([]byte, _ivLen)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	copy(encryptedData[_versionLen+_saltLen:], iv)

	encryptionKey := pbkdf2.Key(key, salt, _pbkdf2c, _pbkdf2KeyLen, sha256.New)

	// Encrypt the data with the first 16 bytes of the encryption key
	aesCipher, err := aes.NewCipher(encryptionKey[:16])
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(aesCipher, iv)
	stream.XORKeyStream(encryptedData[_versionLen+_saltLen+_ivLen+_checksumLen:], data)

	// Generate the checksum
	h := sha256.New()
	if _, err := h.Write(encryptionKey[16:32]); err != nil {
		return nil, err
	}
	if _, err := h.Write(encryptedData[_versionLen+_saltLen+_ivLen+_checksumLen:]); err != nil {
		return nil, err
	}
	checksum := h.Sum(nil)
	copy(encryptedData[_versionLen+_saltLen+_ivLen:_versionLen+_saltLen+_ivLen+_checksumLen], checksum)

	return encryptedData, nil
}
