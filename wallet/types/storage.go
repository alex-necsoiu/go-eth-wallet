package types

import (
	"context"

	"github.com/google/uuid"
)

// Store is the interface for wallet stores.  It is used to store and access data provided by wallets, both wallets themselves
// as well as keys inside the wallets.
type Store interface {
	// Name provides the name of the store
	Name() string

	// StoreWallet stores wallet data.  It will fail if it cannot store the data.
	StoreWallet(walletID uuid.UUID, walletName string, data []byte) error

	// RetrieveWallet retrieves wallet data for all wallets.
	RetrieveWallets() <-chan []byte

	// RetrieveWallet retrieves wallet data for a wallet with a given name.
	// It will fail if it cannot retrieve the data.
	RetrieveWallet(walletName string) ([]byte, error)

	// RetrieveWalletByID retrieves wallet data for a wallet with a given ID.
	// It will fail if it cannot retrieve the data.
	RetrieveWalletByID(walletID uuid.UUID) ([]byte, error)

	// StoreAccount stores account data.  It will fail if it cannot store the data.
	StoreAccount(walletID uuid.UUID, accountID uuid.UUID, data []byte) error

	// RetrieveAccounts retrieves account information for all accounts.
	RetrieveAccounts(walletID uuid.UUID) <-chan []byte

	// RetrieveAccount retrieves account data for a wallet with a given ID.
	// It will fail if it cannot retrieve the data.
	RetrieveAccount(walletID uuid.UUID, accountID uuid.UUID) ([]byte, error)

	// StoreAccountsIndex stores the index of accounts for a given wallet.
	StoreAccountsIndex(walletID uuid.UUID, data []byte) error

	// RetrieveAccountsIndex retrieves the index of accounts for a given wallet.
	RetrieveAccountsIndex(walletID uuid.UUID) ([]byte, error)
}

// BatchStorer is an interface for storing account batches.
type BatchStorer interface {
	// StoreBatch stores wallet batch data.  It will fail if it cannot store the data.
	StoreBatch(ctx context.Context, walletID uuid.UUID, walletName string, data []byte) error
}

// BatchRetriever is an interface for retrieving account batches.
type BatchRetriever interface {
	// RetrieveBatch retrieves the batch of accounts for a given wallet.
	RetrieveBatch(ctx context.Context, walletID uuid.UUID) ([]byte, error)
}

// StoreProvider is the interface provides a store.
type StoreProvider interface {
	// Store returns the store.
	Store() Store
}

// StoreLocationProvider provides the location of the store.
type StoreLocationProvider interface {
	Location() string
}
