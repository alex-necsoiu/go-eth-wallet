package filesystem_test

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	filesystem "github.com/alex-necsoiu/go-eth-wallet/filesystem"
	"github.com/alex-necsoiu/go-eth-wallet/indexer"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestStoreRetrieveIndex(t *testing.T) {
	path := filepath.Join(os.TempDir(), fmt.Sprintf("%s-%d", t.Name(), rand.Int31()))
	defer os.RemoveAll(path)
	store := filesystem.New(filesystem.WithLocation(path))

	walletID := uuid.New()
	walletName := "test wallet"
	walletData := []byte(fmt.Sprintf(`{"name":%q,"uuid":%q}`, walletName, walletID.String()))
	accountID := uuid.New()
	accountName := "test account"
	accountData := []byte(fmt.Sprintf(`{"name":%q,"uuid":%q}`, accountName, accountID.String()))

	index := indexer.New()
	index.Add(accountID, accountName)

	err := store.StoreWallet(walletID, walletName, walletData)
	require.Nil(t, err)
	err = store.StoreAccount(walletID, accountID, accountData)
	require.Nil(t, err)

	serializedIndex, err := index.Serialize()
	require.Nil(t, err)
	err = store.StoreAccountsIndex(walletID, serializedIndex)
	require.Nil(t, err)

	fetchedIndex, err := store.RetrieveAccountsIndex(walletID)
	require.Nil(t, err)

	reIndex, err := indexer.Deserialize(fetchedIndex)
	require.Nil(t, err)

	fetchedAccountName, exists := reIndex.Name(accountID)
	require.Equal(t, true, exists)
	require.Equal(t, accountName, fetchedAccountName)

	fetchedAccountID, exists := reIndex.ID(accountName)
	require.Equal(t, true, exists)
	require.Equal(t, accountID, fetchedAccountID)
}
