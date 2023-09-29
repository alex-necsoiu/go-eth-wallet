package scratch

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// StoreBatch stores wallet batch data.  It will fail if it cannot store the data.
func (s *Store) StoreBatch(_ context.Context, walletID uuid.UUID, _ string, data []byte) error {
	s.batchMu.Lock()
	s.walletBatches[walletID] = data
	s.batchMu.Unlock()

	return nil
}

// RetrieveBatch retrieves the batch of accounts for a given wallet.
func (s *Store) RetrieveBatch(_ context.Context, walletID uuid.UUID) ([]byte, error) {
	s.batchMu.RLock()
	batch, exists := s.walletBatches[walletID]
	s.batchMu.RUnlock()

	if !exists {
		return nil, errors.New("no batch")
	}

	return batch, nil
}
