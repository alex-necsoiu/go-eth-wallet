package hd_test

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"

	hd "github.com/alex-necsoiu/go-eth-wallet/hd"
	keystorev4 "github.com/alex-necsoiu/go-eth-wallet/keystore"
	scratch "github.com/alex-necsoiu/go-eth-wallet/scratch"
	e2wtypes "github.com/alex-necsoiu/go-eth-wallet/wallet/types"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	tests := []struct {
		name              string
		accountName       string
		walletPassphrase  []byte
		accountPassphrase []byte
		err               error
	}{
		{
			name:        "Empty",
			accountName: "",
			err:         errors.New("account name missing"),
		},
		{
			name:        "Invalid",
			accountName: "_bad",
			err:         errors.New(`invalid account name "_bad"`),
		},
		{
			name:        "Good",
			accountName: "test",
		},
		{
			name:        "Duplicate",
			accountName: "test",
			err:         errors.New(`account with name "test" already exists`),
		},
	}

	store := scratch.New()
	encryptor := keystorev4.New()
	seed := []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
	}
	wallet, err := hd.CreateWallet(context.Background(), "test wallet", []byte("wallet passphrase"), store, encryptor, seed)
	require.Nil(t, err)

	// Try to create without unlocking the wallet; should fail.
	_, err = wallet.(e2wtypes.WalletAccountCreator).CreateAccount(context.Background(), "attempt", []byte("test"))
	assert.NotNil(t, err)

	locker, isLocker := wallet.(e2wtypes.WalletLocker)
	require.True(t, isLocker)
	err = locker.Unlock(context.Background(), []byte("wallet passphrase"))
	require.Nil(t, err)
	defer func(t *testing.T) {
		require.NoError(t, locker.Lock(context.Background()))
	}(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err = wallet.(e2wtypes.WalletAccountCreator).CreateAccount(context.Background(), test.accountName, test.accountPassphrase)
			if test.err != nil {
				require.NotNil(t, err)
				assert.Equal(t, test.err.Error(), err.Error())
			} else {
				require.Nil(t, err)
				accountByNameProvider, isAccountByNameProvider := wallet.(e2wtypes.WalletAccountByNameProvider)
				require.True(t, isAccountByNameProvider)
				account, err := accountByNameProvider.AccountByName(context.Background(), test.accountName)
				require.Nil(t, err)
				assert.Equal(t, test.accountName, account.Name())
				pathProvider, isPathProvider := account.(e2wtypes.AccountPathProvider)
				require.True(t, isPathProvider)
				assert.NotNil(t, pathProvider.Path())
				require.Equal(t, wallet.Name(), account.(e2wtypes.AccountWalletProvider).Wallet().Name())

				// Should not be able to obtain private key from a locked account.
				_, err = account.(e2wtypes.AccountPrivateKeyProvider).PrivateKey(context.Background())
				assert.NotNil(t, err)
				locker, isLocker := account.(e2wtypes.AccountLocker)
				require.True(t, isLocker)
				err = locker.Unlock(context.Background(), test.accountPassphrase)
				require.Nil(t, err)
				_, err = account.(e2wtypes.AccountPrivateKeyProvider).PrivateKey(context.Background())
				assert.Nil(t, err)
			}
		})
	}
}

func TestAccountByNameDynamic(t *testing.T) {
	store := scratch.New()
	encryptor := keystorev4.New()
	seed := []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
	}
	wallet, err := hd.CreateWallet(context.Background(), "test wallet", []byte("wallet passphrase"), store, encryptor, seed)
	require.NoError(t, err)

	locker, isLocker := wallet.(e2wtypes.WalletLocker)
	require.True(t, isLocker)
	err = locker.Unlock(context.Background(), []byte("wallet passphrase"))
	require.NoError(t, err)

	accountByNameProvider, isAccountByNameProvider := wallet.(e2wtypes.WalletAccountByNameProvider)
	require.True(t, isAccountByNameProvider)
	account, err := accountByNameProvider.AccountByName(context.Background(), "m/12381/3600/0/0")
	require.NoError(t, err)
	assert.Equal(t,
		[]byte{
			0x94, 0x6e, 0x0f, 0x38, 0xa0, 0x23, 0xb9, 0xf1, 0xad, 0x94, 0x9c, 0xe2, 0xad, 0x85, 0x31, 0xc4,
			0xdb, 0x53, 0x7e, 0x31, 0x34, 0x26, 0x59, 0x9c, 0x2d, 0x9a, 0xe8, 0xab, 0xee, 0xef, 0x7a, 0x43,
			0x3d, 0x06, 0x67, 0x39, 0xf8, 0x16, 0xdd, 0x53, 0x7a, 0xdb, 0x2e, 0x4b, 0x84, 0x11, 0xcc, 0xcb,
		},
		account.PublicKey().Marshal(),
	)
	account, err = accountByNameProvider.AccountByName(context.Background(), "m/12381/3600/1/1/1")
	require.NoError(t, err)
	assert.Equal(t,
		[]byte{
			0x87, 0x27, 0x31, 0x75, 0x58, 0x9b, 0x59, 0x34, 0x41, 0xb3, 0x7d, 0x94, 0x66, 0x4a, 0x88, 0x89,
			0xc5, 0x2a, 0xf5, 0xbb, 0x10, 0x60, 0xca, 0x06, 0x91, 0x27, 0xd4, 0x31, 0x82, 0x12, 0xc4, 0x4f,
			0x1e, 0x2d, 0xdb, 0x77, 0xfa, 0x55, 0xd5, 0x5b, 0x5c, 0xde, 0x58, 0xcc, 0x42, 0x5e, 0xa5, 0xa1,
		},
		account.PublicKey().Marshal(),
	)
}

func TestCreatePathedAccount(t *testing.T) {
	store := scratch.New()
	encryptor := keystorev4.New()
	seed := []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
	}
	wallet, err := hd.CreateWallet(context.Background(), "test wallet", []byte("wallet passphrase"), store, encryptor, seed)
	require.Nil(t, err)
	locker, isLocker := wallet.(e2wtypes.WalletLocker)
	require.True(t, isLocker)
	err = locker.Unlock(context.Background(), []byte("wallet passphrase"))
	require.Nil(t, err)

	// Create an account without a path.
	_, err = wallet.(e2wtypes.WalletAccountCreator).CreateAccount(context.Background(), "Test", []byte("account passphrase"))
	require.Nil(t, err)
	// Attempt to create an account with the same path; should fail.
	_, err = wallet.(e2wtypes.WalletPathedAccountCreator).CreatePathedAccount(context.Background(), "m/12381/3600/0/0", "Test 2", []byte("account passphrase"))
	require.EqualError(t, err, `account with path "m/12381/3600/0/0" already exists`)

	// Attempt to create an account with the a different path; should succeed.
	_, err = wallet.(e2wtypes.WalletPathedAccountCreator).CreatePathedAccount(context.Background(), "m/12381/3600/1/2/3", "Test 3", []byte("account passphrase"))
	require.Nil(t, err)
	// Attempt to create an account with the the same path; should fail.
	_, err = wallet.(e2wtypes.WalletPathedAccountCreator).CreatePathedAccount(context.Background(), "m/12381/3600/1/2/3", "Test 4", []byte("account passphrase"))
	require.EqualError(t, err, `account with path "m/12381/3600/1/2/3" already exists`)

	// Attempt to create an account with the highest legal index; should succeed.
	_, err = wallet.(e2wtypes.WalletPathedAccountCreator).CreatePathedAccount(context.Background(), "m/12381/3600/1/2/4294967295", "Test 5", []byte("account passphrase"))
	require.Nil(t, err)
	// Attempt to create an account with the lowest illegal index; should fail.
	_, err = wallet.(e2wtypes.WalletPathedAccountCreator).CreatePathedAccount(context.Background(), "m/12381/3600/1/2/4294967296", "Test 6", []byte("account passphrase"))
	require.EqualError(t, err, `failed to create private key for account "Test 6": invalid index "4294967296" at path component 5`)
}

func TestCreatePathedAccountConflict(t *testing.T) {
	store := scratch.New()
	encryptor := keystorev4.New()
	seed := []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
	}
	wallet, err := hd.CreateWallet(context.Background(), "test wallet", []byte("wallet passphrase"), store, encryptor, seed)
	require.Nil(t, err)
	locker, isLocker := wallet.(e2wtypes.WalletLocker)
	require.True(t, isLocker)
	err = locker.Unlock(context.Background(), []byte("wallet passphrase"))
	require.Nil(t, err)

	// Create an account with the explicit path of the first index.
	_, err = wallet.(e2wtypes.WalletPathedAccountCreator).CreatePathedAccount(context.Background(), "m/12381/3600/0/0", "Test 1", []byte("account passphrase"))
	require.Nil(t, err)

	// Now create an unpathed account; should have the next index.
	account, err := wallet.(e2wtypes.WalletAccountCreator).CreateAccount(context.Background(), "Test 2", []byte("account passphrase"))
	require.Nil(t, err)
	require.Equal(t, "m/12381/3600/1/0", account.(e2wtypes.AccountPathProvider).Path())

	// Now create another unpathed account; should have the next index.
	account, err = wallet.(e2wtypes.WalletAccountCreator).CreateAccount(context.Background(), "Test 3", []byte("account passphrase"))
	require.Nil(t, err)
	require.Equal(t, "m/12381/3600/2/0", account.(e2wtypes.AccountPathProvider).Path())
}

func TestConcurrentCreate(t *testing.T) {
	store := scratch.New()
	encryptor := keystorev4.New()
	seed := []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
	}
	wallet, err := hd.CreateWallet(context.Background(), "test wallet", []byte("wallet passphrase"), store, encryptor, seed)
	require.NoError(t, err)
	locker, isLocker := wallet.(e2wtypes.WalletLocker)
	require.True(t, isLocker)
	require.NoError(t, locker.Unlock(context.Background(), []byte("wallet passphrase")))

	// Create a number of runners that will try to create accounts simultaneously.
	var runWG sync.WaitGroup
	var setupWG sync.WaitGroup
	starter := make(chan any)
	numAccounts := 64
	for i := 0; i < numAccounts; i++ {
		setupWG.Add(1)
		runWG.Add(1)
		go func() {
			id := rand.Uint32()
			name := fmt.Sprintf("Test account %d", id)
			setupWG.Done()

			<-starter

			_, err := wallet.(e2wtypes.WalletAccountCreator).CreateAccount(context.Background(), name, []byte("test"))
			require.NoError(t, err)
			runWG.Done()
		}()
	}

	// Wait for setup to complete.
	setupWG.Wait()

	// Start the jobs by closing the channel.
	close(starter)

	// Wait for run to complete
	runWG.Wait()

	// Confirm that all accounts have been created with different paths.
	wallet, err = hd.OpenWallet(context.Background(), "test wallet", store, encryptor)
	require.NoError(t, err)
	paths := make(map[string]struct{})

	for account := range wallet.Accounts(context.Background()) {
		paths[account.(e2wtypes.AccountPathProvider).Path()] = struct{}{}
	}
	require.Equal(t, numAccounts, len(paths))
}

func TestAccountLocking(t *testing.T) {
	store := scratch.New()
	encryptor := keystorev4.New()
	seed := []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
	}
	wallet, err := hd.CreateWallet(context.Background(), "test wallet", []byte("wallet passphrase"), store, encryptor, seed)
	require.NoError(t, err)
	require.NoError(t, wallet.(e2wtypes.WalletLocker).Unlock(context.Background(), []byte("wallet passphrase")))

	account, err := wallet.(e2wtypes.WalletAccountCreator).CreateAccount(context.Background(), "test account", []byte("account passphrase"))
	require.NoError(t, err)

	locker, isLocker := account.(e2wtypes.AccountLocker)
	require.True(t, isLocker)

	// Ensure the wallet is not unlocked to begin with.
	unlocked, err := locker.IsUnlocked(context.Background())
	require.NoError(t, err)
	require.False(t, unlocked)

	// Ensure the wallet is not unlocked when an incorrect passphrase is supplied.
	require.EqualError(t, locker.Unlock(context.Background(), []byte("bad passphrase")), "incorrect passphrase")
	unlocked, err = locker.IsUnlocked(context.Background())
	require.NoError(t, err)
	require.False(t, unlocked)

	// Ensure the wallet is unlocked when the correct passphrase is supplied.
	require.NoError(t, locker.Unlock(context.Background(), []byte("account passphrase")))
	unlocked, err = locker.IsUnlocked(context.Background())
	require.NoError(t, err)
	require.True(t, unlocked)

	// ensure the wallet remains unlocked when already unlocked.
	require.NoError(t, locker.Unlock(context.Background(), []byte("account passphrase")))
	unlocked, err = locker.IsUnlocked(context.Background())
	require.NoError(t, err)
	require.True(t, unlocked)
}
