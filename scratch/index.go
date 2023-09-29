package scratch

import (
	"errors"

	"github.com/google/uuid"
)

// StoreAccountsIndex stores the account index.
func (s *Store) StoreAccountsIndex(walletID uuid.UUID, data []byte) error {
	s.accountIndex[walletID] = data
	return nil
}

// RetrieveAccountsIndex retrieves the account index.
func (s *Store) RetrieveAccountsIndex(walletID uuid.UUID) ([]byte, error) {
	index, exists := s.accountIndex[walletID]
	if !exists {
		return nil, errors.New("not found")
	}

	return index, nil
}
