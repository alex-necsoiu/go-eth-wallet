package scratch

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
)

// StoreAccount stores an account.  It will fail if it cannot store the data.
// Note this will overwrite an existing account with the same ID.  It will not, however, allow multiple accounts with the same
// name to co-exist in the same wallet.
func (s *Store) StoreAccount(walletID uuid.UUID, accountID uuid.UUID, data []byte) error {
	s.accountMu.Lock()
	s.accounts[walletID.String()][accountID.String()] = data
	s.accountMu.Unlock()

	return nil
}

// RetrieveAccount retrieves account-level data.  It will fail if it cannot retrieve the data.
func (s *Store) RetrieveAccount(walletID uuid.UUID, accountID uuid.UUID) ([]byte, error) {
	for data := range s.RetrieveAccounts(walletID) {
		info := &struct {
			ID uuid.UUID `json:"uuid"`
		}{}
		err := json.Unmarshal(data, info)
		if err == nil && info.ID == accountID {
			return data, nil
		}
	}

	return nil, errors.New("account not found")
}

// RetrieveAccounts retrieves all account-level data for a wallet.
func (s *Store) RetrieveAccounts(walletID uuid.UUID) <-chan []byte {
	ch := make(chan []byte, 1024)
	go func() {
		s.accountMu.RLock()
		for _, account := range s.accounts[walletID.String()] {
			ch <- account
		}
		s.accountMu.RUnlock()
		close(ch)
	}()

	return ch
}
