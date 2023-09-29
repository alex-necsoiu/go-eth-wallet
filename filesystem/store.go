// Package filesystem is an Ethereum wallet store on a local filesystem.
package filesystem

import (
	wtypes "github.com/alex-necsoiu/go-eth-wallet/wallet/types"
	"github.com/shibukawa/configdir"
)

// options are the options for the filesystem store.
type options struct {
	passphrase []byte
	location   string
}

// Option gives options to New.
type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

// WithPassphrase sets the encryption for the store.
func WithPassphrase(passphrase []byte) Option {
	return optionFunc(func(o *options) {
		o.passphrase = passphrase
	})
}

// WithLocation sets the on-filesystem location for the store.
func WithLocation(b string) Option {
	return optionFunc(func(o *options) {
		o.location = b
	})
}

// Store is the store for the wallet.
type Store struct {
	location   string
	passphrase []byte
}

func defaultLocation() string {
	configDirs := configdir.New("ethereum2", "wallets")
	return configDirs.QueryFolders(configdir.Global)[0].Path
}

// New creates a new filesystem store.
// If the path is not supplied a default path is used.
func New(opts ...Option) wtypes.Store {
	options := options{
		location: defaultLocation(),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	return &Store{
		location:   options.location,
		passphrase: options.passphrase,
	}
}

// Name returns the name of this store.
func (s *Store) Name() string {
	return "filesystem"
}

// Location returns the location of this store.
func (s *Store) Location() string {
	return s.location
}
