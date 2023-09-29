package wallet

import (
	"errors"
	"fmt"

	filesystem "github.com/alex-necsoiu/go-eth-wallet/filesystem"
	scratch "github.com/alex-necsoiu/go-eth-wallet/scratch"
	s3 "github.com/alex-necsoiu/go-eth-wallet/store-s3"
	wtypes "github.com/alex-necsoiu/go-eth-wallet/wallet/types"
)

var store wtypes.Store

func init() {
	// Default store is filesystem.
	store = filesystem.New()
}

// SetStore sets a store to use given its name and optional passphrase.
// This does not allow access to all advanced features of stores.  To access these create the stores yourself and set them with
// `UseStore()`.
func SetStore(name string, passphrase []byte) error {
	var store wtypes.Store
	var err error
	switch name {
	case "s3":
		store, err = s3.New(s3.WithPassphrase(passphrase))
	case "filesystem":
		store = filesystem.New(filesystem.WithPassphrase(passphrase))
	case "scratch":
		store = scratch.New()
	default:
		err = fmt.Errorf("unknown wallet store %q", name)
	}
	if err != nil {
		return err
	}

	return UseStore(store)
}

// UseStore sets a store to use.
func UseStore(s wtypes.Store) error {
	if s == nil {
		return errors.New("no store supplied")
	}
	store = s

	return nil
}

// GetStore returns the name of the current store.
func GetStore() string {
	return store.Name()
}
