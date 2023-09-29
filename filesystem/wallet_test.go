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

func TestStoreRetrieveWallet(t *testing.T) {
	path := filepath.Join(os.TempDir(), fmt.Sprintf("TestStoreRetrieveWallet-%d", rand.Int31()))
	defer os.RemoveAll(path)
	store := filesystem.New(filesystem.WithLocation(path))

	walletID := uuid.New()
	walletName := "test wallet"
	data := []byte(fmt.Sprintf(`{"uuid":%q,"name":%q}`, walletID, walletName))

	err := store.StoreWallet(walletID, walletName, data)
	require.Nil(t, err)
	retData, err := store.RetrieveWallet(walletName)
	require.Nil(t, err)
	assert.Equal(t, data, retData)

	for range store.RetrieveWallets() {
	}
}

func TestRetrieveNonExistentWallet(t *testing.T) {
	path := filepath.Join(os.TempDir(), fmt.Sprintf("TestRetrieveNonExistentWallet-%d", rand.Int31()))
	defer os.RemoveAll(path)
	store := filesystem.New(filesystem.WithLocation(path))

	walletName := "test wallet"

	_, err := store.RetrieveWallet(walletName)
	assert.NotNil(t, err)
}
