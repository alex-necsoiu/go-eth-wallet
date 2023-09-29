package unencrypted_test

import (
	"encoding/json"
	"testing"

	unencrypted "github.com/alex-necsoiu/go-eth-wallet/unencrypted"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

func TestRoundTrip(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		passphrase string
		secret     []byte
		err        error
	}{
		{
			name:  "Good",
			input: `{"key":"0x25295f0d1d592a90b333e26e85149708208e9f8e8bc18f6c77bd62f8ad7a6866"}`,
			secret: []byte{
				0x25, 0x29, 0x5f, 0x0d, 0x1d, 0x59, 0x2a, 0x90, 0xb3, 0x33, 0xe2, 0x6e, 0x85, 0x14, 0x97, 0x08,
				0x20, 0x8e, 0x9f, 0x8e, 0x8b, 0xc1, 0x8f, 0x6c, 0x77, 0xbd, 0x62, 0xf8, 0xad, 0x7a, 0x68, 0x66,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			encryptor := unencrypted.New()
			input := make(map[string]any)
			err := json.Unmarshal([]byte(test.input), &input)
			require.Nil(t, err)
			secret, err := encryptor.Decrypt(input, test.passphrase)
			if test.err != nil {
				require.NotNil(t, err)
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, test.secret, secret)
				newInput, err := encryptor.Encrypt(secret, test.passphrase)
				require.Nil(t, err)
				newSecret, err := encryptor.Decrypt(newInput, test.passphrase)
				require.Nil(t, err)
				require.Equal(t, test.secret, newSecret)
			}
		})
	}
}

func TestNameAndVersion(t *testing.T) {
	encryptor := unencrypted.New()
	assert.Equal(t, "unencrypted", encryptor.Name())
	assert.Equal(t, uint(1), encryptor.Version())
}
