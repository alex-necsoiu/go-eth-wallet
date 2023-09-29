package s3_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	s3 "github.com/alex-necsoiu/go-eth-wallet/store-s3"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStoreRetrieveEncryptedWallet(t *testing.T) {
	//nolint:gosec
	rand := rand.New(rand.NewSource(time.Now().Unix()))
	id := fmt.Sprintf("%s-%d", t.Name(), rand.Int31())
	store, err := s3.New(s3.WithID([]byte(id)),
		s3.WithPassphrase([]byte("test")),
		s3.WithCredentialsID(os.Getenv("S3_CREDENTIALS_ID")),
		s3.WithCredentialsSecret(os.Getenv("S3_CREDENTIALS_SECRET")),
		s3.WithBucket(os.Getenv("S3_BUCKET")),
	)
	if err != nil {
		t.Skip("unable to access S3; skipping test")
	}

	walletID := uuid.New()
	walletName := fmt.Sprintf("test wallet for TestStoreRetrieveEncryptedWallet %d", time.Now().UnixNano())
	data := []byte(fmt.Sprintf(`{"uuid":%q,"name":%q}`, walletID, walletName))

	err = store.StoreWallet(walletID, walletName, data)
	require.Nil(t, err)
	retData, err := store.RetrieveWallet(walletName)
	require.Nil(t, err)
	assert.Equal(t, data, retData)

	wallets := false
	for range store.RetrieveWallets() {
		wallets = true
	}
	assert.True(t, wallets)

	store.RetrieveWallets()
}

func TestStoreRetrieveEncryptedAccount(t *testing.T) {
	//nolint:gosec
	rand := rand.New(rand.NewSource(time.Now().Unix()))
	id := fmt.Sprintf("%s-%d", t.Name(), rand.Int31())
	store, err := s3.New(s3.WithID([]byte(id)),
		s3.WithPassphrase([]byte("test")),
		s3.WithCredentialsID(os.Getenv("S3_CREDENTIALS_ID")),
		s3.WithCredentialsSecret(os.Getenv("S3_CREDENTIALS_SECRET")),
		s3.WithBucket(os.Getenv("S3_BUCKET")),
	)
	if err != nil {
		t.Skip("unable to access S3; skipping test")
	}

	walletID := uuid.New()
	walletName := fmt.Sprintf("test wallet for TestStoreRetrieveEncryptedAccount %d", time.Now().UnixNano())
	walletData := []byte(fmt.Sprintf(`{"name":%q,"uuid":%q}`, walletName, walletID.String()))
	accountID := uuid.New()
	accountName := "test account"
	accountData := []byte(fmt.Sprintf(`{"name":%q,"uuid":%q}`, accountName, accountID.String()))

	err = store.StoreWallet(walletID, walletName, walletData)
	require.Nil(t, err)

	err = store.StoreAccount(walletID, accountID, accountData)
	require.Nil(t, err)
	retData, err := store.RetrieveAccount(walletID, accountID)
	require.Nil(t, err)
	require.Equal(t, accountData, retData)

	accounts := false
	for range store.RetrieveAccounts(walletID) {
		accounts = true
	}
	assert.True(t, accounts)
}

func TestBadWalletKey(t *testing.T) {
	//nolint:gosec
	rand := rand.New(rand.NewSource(time.Now().Unix()))
	id := fmt.Sprintf("%s-%d", t.Name(), rand.Int31())
	store, err := s3.New(s3.WithID([]byte(id)),
		s3.WithPassphrase([]byte("test")),
		s3.WithCredentialsID(os.Getenv("S3_CREDENTIALS_ID")),
		s3.WithCredentialsSecret(os.Getenv("S3_CREDENTIALS_SECRET")),
		s3.WithBucket(os.Getenv("S3_BUCKET")),
	)
	if err != nil {
		t.Skip("unable to access S3; skipping test")
	}

	walletID := uuid.New()
	walletName := fmt.Sprintf("test wallet for TestBadWalletKey %d", time.Now().UnixNano())
	data := []byte(fmt.Sprintf(`{"uuid":%q,"name":%q}`, walletID, walletName))

	err = store.StoreWallet(walletID, walletName, data)
	require.Nil(t, err)

	// Open wallet with store with different key; should fail
	store, err = s3.New(s3.WithID([]byte(id)),
		s3.WithPassphrase([]byte("badkey")),
		s3.WithCredentialsID(os.Getenv("S3_CREDENTIALS_ID")),
		s3.WithCredentialsSecret(os.Getenv("S3_CREDENTIALS_SECRET")),
		s3.WithBucket(os.Getenv("S3_BUCKET")),
	)
	require.Nil(t, err)
	_, err = store.RetrieveWallet(walletName)
	require.NotNil(t, err)
}
