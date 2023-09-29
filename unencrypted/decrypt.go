package unencrypted

import (
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

// Decrypt decrypts the data provided, returning the secret.
func (e *Encryptor) Decrypt(data map[string]any, _ string) ([]byte, error) {
	if data == nil {
		return nil, errors.New("no data supplied")
	}
	// Marshal the map and unmarshal it back in to a format we can work with.
	b, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse keystore")
	}
	ks := &unencrypted{}
	err = json.Unmarshal(b, &ks)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse keystore")
	}

	if ks.Key == "" {
		return nil, errors.New("key missing")
	}

	key, err := hex.DecodeString(strings.TrimPrefix(ks.Key, "0x"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode key")
	}

	return key, nil
}
