package s3

import (
	"github.com/google/uuid"
)

func (s *Store) walletPath(walletID uuid.UUID) string {
	return join(s.path, walletID.String())
}

func (s *Store) walletHeaderPath(walletID uuid.UUID) string {
	return join(s.walletPath(walletID), walletID.String())
}

func (s *Store) accountPath(walletID uuid.UUID, accountID uuid.UUID) string {
	return join(s.walletPath(walletID), accountID.String())
}

func (s *Store) walletIndexPath(walletID uuid.UUID) string {
	return join(s.walletPath(walletID), "index")
}

func (s *Store) walletBatchPath(walletID uuid.UUID) string {
	return join(s.walletPath(walletID), "batch")
}

// join joins multiple segments of a path.
func join(elem ...string) string {
	res := ""
	for _, e := range elem {
		if e == "" {
			continue
		}
		if res != "" {
			res += "/"
		}
		res += e
	}

	return res
}
