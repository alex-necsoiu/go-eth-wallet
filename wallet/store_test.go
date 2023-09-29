package wallet_test

import (
	"testing"

	scratch "github.com/alex-necsoiu/go-eth-wallet/scratch"
	wallet "github.com/alex-necsoiu/go-eth-wallet/wallet"
	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {
	// Ensure default store is set.
	require.Equal(t, "filesystem", wallet.GetStore())

	// Attempt to set a nil store; should error.
	require.EqualError(t, wallet.UseStore(nil), "no store supplied")

	// Attempt to set a different store.
	require.NoError(t, wallet.UseStore(scratch.New()))

	// Confirm the store has been set.
	require.Equal(t, "scratch", wallet.GetStore())

	// Attempt to switch stores.
	require.NoError(t, wallet.SetStore("filesystem", nil))
	require.Equal(t, "filesystem", wallet.GetStore())
	require.NoError(t, wallet.SetStore("scratch", nil))
	require.Equal(t, "scratch", wallet.GetStore())
	require.EqualError(t, wallet.SetStore("unknown", nil), "unknown wallet store \"unknown\"")
}
