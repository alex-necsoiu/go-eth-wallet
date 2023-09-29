// Package types provides generic types for the Ethereum consensus system.
package types

import (
	bls "github.com/herumi/bls-eth-go-binary/bls"
)

// InitBLS initialises the BLS library with the appropriate curve and parameters for Ethereum 2.
func InitBLS() error {
	if err := bls.Init(bls.BLS12_381); err != nil {
		return err
	}
	return bls.SetETHmode(bls.EthModeDraft07)
}
