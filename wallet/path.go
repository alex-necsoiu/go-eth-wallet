package wallet

import (
	"errors"
	"strings"
)

// WalletAndAccountNames breaks an account in to wallet and account names.
//
//nolint:revive
func WalletAndAccountNames(account string) (string, string, error) {
	if len(account) == 0 {
		return "", "", errors.New("invalid account format")
	}
	index := strings.Index(account, "/")
	if index == -1 {
		// Just the wallet
		return account, "", nil
	}
	if index == 0 {
		return "", "", errors.New("invalid account format")
	}
	if index == len(account)-1 {
		// Trailing /
		return account[:index], "", nil
	}

	return account[:index], account[index+1:], nil
}
