package filesystem

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncryptIfRequired(t *testing.T) {
	path := filepath.Join(os.TempDir(), fmt.Sprintf("%s-%d", t.Name(), rand.Int31()))
	defer os.RemoveAll(path)

	tests := []struct {
		name  string
		store *Store
		data  []byte
		err   string
	}{
		{
			name:  "NoData",
			store: &Store{},
		},
		{
			name:  "NoPassphrase",
			store: &Store{},
			data:  []byte(`{"test":true}`),
		},
		{
			name: "ShortData",
			data: []byte(`{"test":true}`),
			store: &Store{
				passphrase: []byte("test passphrase"),
			},
			err: "data must be at least 16 bytes",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := test.store.encryptIfRequired(test.data)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDecryptIfRequired(t *testing.T) {
	path := filepath.Join(os.TempDir(), fmt.Sprintf("%s-%d", t.Name(), rand.Int31()))
	defer os.RemoveAll(path)

	tests := []struct {
		name  string
		store *Store
		data  []byte
		err   string
	}{
		{
			name:  "NoData",
			store: &Store{},
		},
		{
			name:  "NoPassphrase",
			store: &Store{},
			data:  []byte(`{"test":true}`),
		},
		{
			name: "ShortData",
			data: []byte(`{"test":true}`),
			store: &Store{
				passphrase: []byte("test passphrase"),
			},
			err: "data must be at least 16 bytes",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := test.store.decryptIfRequired(test.data)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
