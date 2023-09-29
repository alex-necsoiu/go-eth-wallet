package filesystem_test

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	filesystem "github.com/alex-necsoiu/go-eth-wallet/filesystem"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStoreRetrieveAccount(t *testing.T) {
	path := filepath.Join(os.TempDir(), fmt.Sprintf("%s-%d", t.Name(), rand.Int31()))
	defer os.RemoveAll(path)
	store := filesystem.New(filesystem.WithLocation(path))

	walletID := uuid.New()
	walletName := "test wallet"
	walletData := []byte(fmt.Sprintf(`{"name":%q,"uuid":%q}`, walletName, walletID.String()))
	accountID := uuid.New()
	accountName := "test account"
	accountData := []byte(fmt.Sprintf(`{"name":%q,"uuid":%q}`, accountName, accountID.String()))

	err := store.StoreWallet(walletID, walletName, walletData)
	require.Nil(t, err)
	err = store.StoreAccount(walletID, accountID, accountData)
	require.Nil(t, err)
	retData, err := store.RetrieveAccount(walletID, accountID)
	require.Nil(t, err)
	assert.Equal(t, accountData, retData)

	store.RetrieveWallets()

	_, err = store.RetrieveAccount(walletID, uuid.New())
	assert.NotNil(t, err)
}

func TestDuplicateAccounts(t *testing.T) {
	path := filepath.Join(os.TempDir(), fmt.Sprintf("%s-%d", t.Name(), rand.Int31()))
	defer os.RemoveAll(path)
	store := filesystem.New(filesystem.WithLocation(path))

	walletID := uuid.New()
	walletName := "test wallet"
	walletData := []byte(fmt.Sprintf(`{"name":%q,"uuid":%q}`, walletName, walletID.String()))
	accountID := uuid.New()
	accountName := "test account"
	accountData := []byte(fmt.Sprintf(`{"name":%q,"uuid":%q}`, accountName, accountID.String()))

	err := store.StoreWallet(walletID, walletName, walletData)
	require.Nil(t, err)
	err = store.StoreAccount(walletID, accountID, accountData)
	require.Nil(t, err)
}

func TestRetrieveNonExistentAccount(t *testing.T) {
	path := filepath.Join(os.TempDir(), fmt.Sprintf("TestRetrieveNonExistentAccount-%d", rand.Int31()))
	defer os.RemoveAll(path)
	store := filesystem.New(filesystem.WithLocation(path))

	walletID := uuid.New()
	accountID := uuid.New()

	_, err := store.RetrieveAccount(walletID, accountID)
	assert.NotNil(t, err)
}

func TestStoreNonExistentAccount(t *testing.T) {
	path := filepath.Join(os.TempDir(), fmt.Sprintf("TestStoreNonExistentAccount-%d", rand.Int31()))
	defer os.RemoveAll(path)
	store := filesystem.New(filesystem.WithLocation(path))

	walletID := uuid.New()
	accountID := uuid.New()
	data := []byte(fmt.Sprintf(`"uuid":%q,"name":"test account"}`, accountID))

	err := store.StoreAccount(walletID, accountID, data)
	assert.NotNil(t, err)
}
