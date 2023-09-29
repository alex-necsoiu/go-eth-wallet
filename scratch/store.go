package scratch

import (
	"sync"

	wtypes "github.com/alex-necsoiu/go-eth-wallet/wallet/types"
	"github.com/google/uuid"
)

// Store is the store for the wallet.
type Store struct {
	wallets       map[string][]byte
	walletMu      sync.RWMutex
	accounts      map[string]map[string][]byte
	accountMu     sync.RWMutex
	accountIndex  map[uuid.UUID][]byte
	batchMu       sync.RWMutex
	walletBatches map[uuid.UUID][]byte
}

// New creates a new scratch store.
func New() wtypes.Store {
	return &Store{
		wallets:       make(map[string][]byte),
		accounts:      make(map[string]map[string][]byte),
		accountIndex:  make(map[uuid.UUID][]byte),
		walletBatches: make(map[uuid.UUID][]byte),
	}
}

// Name returns the name of this store.
func (s *Store) Name() string {
	return "scratch"
}
