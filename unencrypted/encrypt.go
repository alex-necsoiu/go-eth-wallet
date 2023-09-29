package unencrypted

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// Encrypt encrypts data.
func (e *Encryptor) Encrypt(secret []byte, _ string) (map[string]any, error) {
	if secret == nil {
		return nil, errors.New("no secret")
	}

	// Build the output.
	output := &unencrypted{
		Key: fmt.Sprintf("%#x", secret),
	}

	// We need to return a generic map; go to JSON and back to obtain it.
	bytes, err := json.Marshal(output)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal output")
	}
	res := make(map[string]any)
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal output")
	}

	return res, nil
}
