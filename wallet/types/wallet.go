// Copyright 2019 - 2023 Weald Technology Trading.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"context"

	"github.com/google/uuid"
)

// WalletIDProvider is the interface for wallets that can provide an ID.
type WalletIDProvider interface {
	// ID provides the ID for the wallet.
	ID() uuid.UUID
}

// WalletNameProvider is the interface for wallets that can provide a name.
type WalletNameProvider interface {
	// Name provides the name for the wallet.
	Name() string
}

// WalletTypeProvider is the interface for wallets that can provide a type.
type WalletTypeProvider interface {
	// Type provides the type for the wallet.
	Type() string
}

// WalletVersionProvider is the interface for wallets that can provide a version.
type WalletVersionProvider interface {
	// Version provides the version of the wallet.
	Version() uint
}

// WalletLocker is the interface for wallets that can be locked and unlocked.
type WalletLocker interface {
	// Lock locks the wallet.  A locked account cannot create new accounts.
	Lock(ctx context.Context) error

	// Unlock unlocks the wallet.  An unlocked account can create new accounts.
	Unlock(ctx context.Context, passphrase []byte) error

	// IsUnlocked returns true if the wallet is unlocked.
	IsUnlocked(ctx context.Context) (bool, error)
}

// WalletAccountsProvider is the interface for wallets that provide account information.
type WalletAccountsProvider interface {
	// Accounts provides all accounts in the wallet.
	Accounts(ctx context.Context) <-chan Account
}

// WalletAccountByIDProvider is the interface for wallets that provide an account given its ID.
type WalletAccountByIDProvider interface {
	// AccountByID provides a single account from the wallet given its ID.
	// This will error if the account is not found.
	AccountByID(ctx context.Context, id uuid.UUID) (Account, error)
}

// WalletAccountByNameProvider is the interface for wallets that provide an account given its name.
type WalletAccountByNameProvider interface {
	// AccountByName provides a single account from the wallet given its name.
	// This will error if the account is not found.
	AccountByName(ctx context.Context, name string) (Account, error)
}

// WalletAccountsByPathProvider is the interface for wallets that provide accounts given a path.
type WalletAccountsByPathProvider interface {
	// AccountsByPath provides all matching accounts in the wallet.
	AccountsByPath(ctx context.Context, path string) <-chan Account
}

// WalletAccountCreator is the interface for wallets that can create accounts.
type WalletAccountCreator interface {
	// CreateAccount creates a new account in the wallet.
	// The only rule for names is that they cannot start with an underscore (_) character.
	// This will error if an account with the name already exists.
	CreateAccount(ctx context.Context, name string, passphrase []byte) (Account, error)
}

// WalletPathedAccountCreator is the interface for wallets that can create accounts with explicit HD paths.
type WalletPathedAccountCreator interface {
	// CreatePathedAccount creates a new account in the wallet with a given path.
	// The only rule for names is that they cannot start with an underscore (_) character.
	// This will error if an account with the name or path already exists.
	CreatePathedAccount(ctx context.Context, path string, name string, passphrase []byte) (Account, error)
}

// WalletDistributedAccountCreator is the interface for wallets that can create distributed accounts.
type WalletDistributedAccountCreator interface {
	// CreateDistributedAccount creates a new distributed account in the wallet.
	// The only rule for names is that they cannot start with an underscore (_) character.
	// This will error if an account with the name already exists.
	CreateDistributedAccount(ctx context.Context,
		name string,
		particpants uint32,
		signingThreshold uint32,
		passphrase []byte,
	) (
		Account,
		error,
	)
}

// WalletExporter is the interface for wallets that can export themselves.
type WalletExporter interface {
	// Export exports the entire wallet, protected by an additional passphrase.
	Export(ctx context.Context, passphrase []byte) ([]byte, error)
}

type WalletBatchCreator interface {
	// BatchWallet encrypts all accounts in a single entity, allowing for faster
	// decryption of wallets with large numbers of accounts.
	BatchWallet(ctx context.Context, passphrases []string, batchPassphrase string) error
}

// WalletAccountImporter is the interface for wallets that can import accounts.
type WalletAccountImporter interface {
	// ImportAccount creates a new account in the wallet from an existing private key.
	// The only rule for names is that they cannot start with an underscore (_) character.
	// This will error if an account with the name already exists.
	ImportAccount(ctx context.Context, name string, key []byte, passphrase []byte) (Account, error)
}

// WalletDistributedAccountImporter is the interface for wallets that can import distributed accounts.
type WalletDistributedAccountImporter interface {
	// ImportDistributedAccount creates a new distributed account in the wallet from provided data.
	// The only rule for names is that they cannot start with an underscore (_) character.
	// This will error if an account with the name already exists.
	ImportDistributedAccount(ctx context.Context,
		name string,
		key []byte,
		signingThreshold uint32,
		verificationVector [][]byte,
		participants map[uint64]string,
		passphrase []byte) (Account, error)
}

// WalletShardedAccountImporter is the interface for wallets that can import sharded accounts.
type WalletShardedAccountImporter interface {
	// ImportShardedAccount creates a new sharded account in the wallet from provided data.
	// The only rule for names is that they cannot start with an underscore (_) character.
	// This will error if an account with the name already exists.
	ImportShardedAccount(ctx context.Context,
		name string,
		key []byte,
		signingThreshold uint32,
		compositePublicKey []byte,
		participants map[uint64]string,
		passphrase []byte) (Account, error)
}

// Wallet is a generic interface for wallets, providing minimal required functionality.
type Wallet interface {
	WalletIDProvider
	WalletTypeProvider
	WalletNameProvider
	WalletVersionProvider
	WalletAccountsProvider
}
