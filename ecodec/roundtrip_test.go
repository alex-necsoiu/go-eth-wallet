package ecodec

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		key  []byte
		err  error
	}{
		{
			name: "Good1",
			data: _byteArray("0102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f20"),
			key:  _byteArray("0102030405060708090a0b0c0d0e0f10"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			encryptedData, err := Encrypt(test.data, test.key)
			if test.err != nil {
				require.NotNil(t, err)
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				require.Nil(t, err)
				data, err := Decrypt(encryptedData, test.key)
				require.Nil(t, err)
				assert.Equal(t, test.data, data)
			}
		})
	}
}
