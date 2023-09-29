package filesystem

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// StoreBatch stores wallet batch data.  It will fail if it cannot store the data.
func (s *Store) StoreBatch(_ context.Context, walletID uuid.UUID, _ string, data []byte) error {
	// Ensure wallet exists.
	_, err := s.RetrieveWalletByID(walletID)
	if err != nil {
		return err
	}

	data, err = s.encryptIfRequired(data)
	if err != nil {
		return errors.Wrap(err, "failed to encrypt batch")
	}

	path := s.walletBatchPath(walletID)

	return os.WriteFile(path, data, 0o600)
}

// RetrieveBatch retrieves the batch of accounts for a given wallet.
func (s *Store) RetrieveBatch(_ context.Context, walletID uuid.UUID) ([]byte, error) {
	// Ensure wallet exists.
	_, err := s.RetrieveWalletByID(walletID)
	if err != nil {
		return nil, err
	}

	path := s.walletBatchPath(walletID)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read batch")
	}

	return s.decryptIfRequired(data)
}
