package filesystem

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func (s *Store) walletPath(walletID uuid.UUID) string {
	return filepath.FromSlash(filepath.Join(s.location, walletID.String()))
}

func (s *Store) walletHeaderPath(walletID uuid.UUID) string {
	return filepath.FromSlash(filepath.Join(s.location, walletID.String(), walletID.String()))
}

func (s *Store) accountPath(walletID uuid.UUID, accountID uuid.UUID) string {
	return filepath.FromSlash(filepath.Join(s.location, walletID.String(), accountID.String()))
}

func (s *Store) walletIndexPath(walletID uuid.UUID) string {
	return filepath.FromSlash(filepath.Join(s.walletPath(walletID), "index"))
}

func (s *Store) walletBatchPath(walletID uuid.UUID) string {
	return filepath.FromSlash(filepath.Join(s.walletPath(walletID), "batch"))
}

func (s *Store) ensureWalletPathExists(walletID uuid.UUID) error {
	path := s.walletPath(walletID)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, 0o700)
		if err != nil {
			return fmt.Errorf("failed to create wallet directory at %s", path)
		}
	}

	return nil
}
