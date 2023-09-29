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

func TestWalletInterfaces(t *testing.T) {
	store := scratch.New()
	encryptor := keystorev4.New()
	wallet, err := nd.CreateWallet(context.Background(), "test wallet", store, encryptor)
	assert.Nil(t, err)

	_, isWalletIDProvider := wallet.(e2wtypes.WalletIDProvider)
	assert.True(t, isWalletIDProvider)
	_, isWalletNameProvider := wallet.(e2wtypes.WalletNameProvider)
	assert.True(t, isWalletNameProvider)
	_, isWalletTypeProvider := wallet.(e2wtypes.WalletTypeProvider)
	assert.True(t, isWalletTypeProvider)
	_, isWalletVersionProvider := wallet.(e2wtypes.WalletVersionProvider)
	assert.True(t, isWalletVersionProvider)
	_, isWalletLocker := wallet.(e2wtypes.WalletLocker)
	assert.True(t, isWalletLocker)
	_, isWalletAccountsProvider := wallet.(e2wtypes.WalletAccountsProvider)
	assert.True(t, isWalletAccountsProvider)
	_, isWalletAccountByIDProvider := wallet.(e2wtypes.WalletAccountByIDProvider)
	assert.True(t, isWalletAccountByIDProvider)
	_, isWalletAccountByNameProvider := wallet.(e2wtypes.WalletAccountByNameProvider)
	assert.True(t, isWalletAccountByNameProvider)
	_, isWalletAccountCreator := wallet.(e2wtypes.WalletAccountCreator)
	assert.True(t, isWalletAccountCreator)
	_, isWalletExporter := wallet.(e2wtypes.WalletExporter)
	assert.True(t, isWalletExporter)
	_, isWalletAccountImporter := wallet.(e2wtypes.WalletAccountImporter)
	assert.True(t, isWalletAccountImporter)
	_, isStoreProvider := wallet.(e2wtypes.StoreProvider)
	assert.True(t, isStoreProvider)
}

func TestCreateWallet(t *testing.T) {
	store := scratch.New()
	encryptor := keystorev4.New()
	wallet, err := nd.CreateWallet(context.Background(), "test wallet", store, encryptor)
	assert.Nil(t, err)

	assert.Equal(t, "test wallet", wallet.Name())
	assert.Equal(t, uint(1), wallet.Version())
	assert.Equal(t, store.Name(), wallet.(e2wtypes.StoreProvider).Store().Name())

	// Try to create another wallet with the same name; should error
	_, err = nd.CreateWallet(context.Background(), "test wallet", store, encryptor)
	assert.NotNil(t, err)
}

func TestWalletUnlockLock(t *testing.T) {
	store := scratch.New()
	encryptor := keystorev4.New()
	wallet, err := nd.CreateWallet(context.Background(), "test wallet", store, encryptor)
	require.NoError(t, err)

	unlocked, err := wallet.(e2wtypes.WalletLocker).IsUnlocked(context.Background())
	require.NoError(t, err)
	require.False(t, unlocked)

	// Unlock.
	require.NoError(t, wallet.(e2wtypes.WalletLocker).Unlock(context.Background(), nil))
	unlocked, err = wallet.(e2wtypes.WalletLocker).IsUnlocked(context.Background())
	require.NoError(t, err)
	require.True(t, unlocked)
	// Unlock again.
	require.NoError(t, wallet.(e2wtypes.WalletLocker).Unlock(context.Background(), nil))
	unlocked, err = wallet.(e2wtypes.WalletLocker).IsUnlocked(context.Background())
	require.NoError(t, err)
	require.True(t, unlocked)
	// Lock
	require.NoError(t, wallet.(e2wtypes.WalletLocker).Lock(context.Background()))
	unlocked, err = wallet.(e2wtypes.WalletLocker).IsUnlocked(context.Background())
	require.NoError(t, err)
	require.False(t, unlocked)
}
