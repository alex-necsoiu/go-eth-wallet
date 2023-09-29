package filesystem

import (
	"os"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// StoreAccountsIndex stores the account index.
func (s *Store) StoreAccountsIndex(walletID uuid.UUID, data []byte) error {
	// Ensure wallet path exists.
	var err error
	if err = s.ensureWalletPathExists(walletID); err != nil {
		return errors.Wrap(err, "wallet path does not exist")
	}

	// Do not encrypt empty index.
	if len(data) != 2 {
		data, err = s.encryptIfRequired(data)
		if err != nil {
			return errors.Wrap(err, "failed to encrypt index")
		}
	}

	path := s.walletIndexPath(walletID)

	return os.WriteFile(path, data, 0o600)
}

// RetrieveAccountsIndex retrieves the account index.
func (s *Store) RetrieveAccountsIndex(walletID uuid.UUID) ([]byte, error) {
	path := s.walletIndexPath(walletID)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read wallet index")
	}
	// Do not decrypt empty index.
	if len(data) == 2 {
		return data, nil
	}

	return s.decryptIfRequired(data)
}
