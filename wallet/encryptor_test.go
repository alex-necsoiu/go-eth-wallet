package wallet_test

import (
	"testing"

	keystorev4 "github.com/alex-necsoiu/go-eth-wallet/keystore"
	unencrypted "github.com/alex-necsoiu/go-eth-wallet/unencrypted"
	wallet "github.com/alex-necsoiu/go-eth-wallet/wallet"
	"github.com/stretchr/testify/require"
)

func TestEncryptor(t *testing.T) {
	// Ensure default encryptor is set.
	require.Equal(t, "keystore", wallet.GetEncryptor())

	// Attempt to set a nil encryptor; should error.
	require.EqualError(t, wallet.UseEncryptor(nil), "no encryptor supplied")

	// Attempt to set a different encryptor.
	require.NoError(t, wallet.UseEncryptor(unencrypted.New()))

	// Confirm the encryptor has been set.
	require.Equal(t, "unencrypted", wallet.GetEncryptor())

	// Attempt to set a different encryptor.
	require.NoError(t, wallet.UseEncryptor(keystorev4.New()))

	// Confirm the encryptor has been set.
	require.Equal(t, "keystore", wallet.GetEncryptor())
}
