package nd_test

import (
	"context"
	"testing"

	keystorev4 "github.com/alex-necsoiu/go-eth-wallet/keystore"
	nd "github.com/alex-necsoiu/go-eth-wallet/nd"
	scratch "github.com/alex-necsoiu/go-eth-wallet/scratch"
	e2wtypes "github.com/alex-necsoiu/go-eth-wallet/wallet/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExportWallet(t *testing.T) {
	store := scratch.New()
	encryptor := keystorev4.New()
	wallet, err := nd.CreateWallet(context.Background(), "test wallet", store, encryptor)
	require.Nil(t, err)
	locker, isLocker := wallet.(e2wtypes.WalletLocker)
	require.True(t, isLocker)
	err = locker.Unlock(context.Background(), []byte{})
	require.Nil(t, err)

	account1, err := wallet.(e2wtypes.WalletAccountCreator).CreateAccount(context.Background(), "Account 1", []byte{})
	require.Nil(t, err)
	account2, err := wallet.(e2wtypes.WalletAccountCreator).CreateAccount(context.Background(), "Account 2", []byte{})
	require.Nil(t, err)

	dump, err := wallet.(e2wtypes.WalletExporter).Export(context.Background(), []byte("dump"))
	require.Nil(t, err)

	// Import it
	store2 := scratch.New()
	wallet2, err := nd.Import(context.Background(), dump, []byte("dump"), store2, encryptor)
	require.Nil(t, err)

	// Confirm the accounts are present
	account1Present := false
	account2Present := false
	for account := range wallet2.Accounts(context.Background()) {
		if account.ID().String() == account1.ID().String() {
			account1Present = true
		}
		if account.ID().String() == account2.ID().String() {
			account2Present = true
		}
	}
	assert.True(t, account1Present && account2Present)

	// Try to import it again; should fail
	_, err = nd.Import(context.Background(), dump, []byte("dump"), store2, encryptor)
	assert.NotNil(t, err)
}
