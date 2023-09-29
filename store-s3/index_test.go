package s3_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/alex-necsoiu/go-eth-wallet/indexer"
	s3 "github.com/alex-necsoiu/go-eth-wallet/store-s3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestStoreRetrieveIndex(t *testing.T) {
	store, err := s3.New(
		s3.WithCredentialsID(os.Getenv("S3_CREDENTIALS_ID")),
		s3.WithCredentialsSecret(os.Getenv("S3_CREDENTIALS_SECRET")),
		s3.WithBucket(os.Getenv("S3_BUCKET")),
	)
	if err != nil {
		t.Skip("unable to access S3; skipping test")
	}

	walletID := uuid.New()
	walletName := fmt.Sprintf("test wallet for TestStoreRetrieveIndex %d", time.Now().UnixNano())
	walletData := []byte(fmt.Sprintf(`{"name":%q,"uuid":%q}`, walletName, walletID.String()))
	accountID := uuid.New()
	accountName := "test account"
	accountData := []byte(fmt.Sprintf(`{"name":%q,"uuid":%q}`, accountName, accountID.String()))

	index := indexer.New()
	index.Add(accountID, accountName)

	err = store.StoreWallet(walletID, walletName, walletData)
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
