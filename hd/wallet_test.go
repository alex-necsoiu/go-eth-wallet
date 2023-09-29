package hd_test

import (
	"context"
	"testing"

	hd "github.com/alex-necsoiu/go-eth-wallet/hd"
	keystorev4 "github.com/alex-necsoiu/go-eth-wallet/keystore"
	scratch "github.com/alex-necsoiu/go-eth-wallet/scratch"
	e2wtypes "github.com/alex-necsoiu/go-eth-wallet/wallet/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInterfaces(t *testing.T) {
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
	_, isWalletPathedAccountCreator := wallet.(e2wtypes.WalletPathedAccountCreator)
	assert.True(t, isWalletPathedAccountCreator)
	_, isWalletExporter := wallet.(e2wtypes.WalletExporter)
	assert.True(t, isWalletExporter)
}

func TestCreateWallet(t *testing.T) {
	tests := []struct {
		name string
		seed []byte
		err  string
	}{
		{
			name: "NoSeed",
			err:  "seed must be 64 bytes",
		},
		{
			name: "ShortSeed",
			seed: []byte{
				0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
				0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
				0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
				0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e,
			},
			err: "seed must be 64 bytes",
		},
		{
			name: "LongSeed",
			seed: []byte{
				0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
				0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
				0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
				0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
				0x40,
			},
			err: "seed must be 64 bytes",
		},
		{
			name: "Good",
			seed: []byte{
				0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
				0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
				0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
				0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
			},
		},
		{
			name: "Dup",
			seed: []byte{
				0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
				0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
				0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
				0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
			},
		},
		{
			name: "Dup",
			seed: []byte{
				0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
				0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
				0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
				0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
			},
			err: "wallet \"Dup\" already exists",
		},
	}

	store := scratch.New()
	encryptor := keystorev4.New()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := hd.CreateWallet(context.Background(), test.name, []byte("wallet passphrase"), store, encryptor, test.seed)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestWalletUnlockLock(t *testing.T) {
	store := scratch.New()
	encryptor := keystorev4.New()
	seed := []byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
	}
	wallet, err := hd.CreateWallet(context.Background(), "test wallet", []byte("pass"), store, encryptor, seed)
	require.NoError(t, err)

	unlocked, err := wallet.(e2wtypes.WalletLocker).IsUnlocked(context.Background())
	require.NoError(t, err)
	require.False(t, unlocked)

	// Unlock.
	require.NoError(t, wallet.(e2wtypes.WalletLocker).Unlock(context.Background(), []byte("pass")))
	unlocked, err = wallet.(e2wtypes.WalletLocker).IsUnlocked(context.Background())
	require.NoError(t, err)
	require.True(t, unlocked)
	// Unlock again.
	require.NoError(t, wallet.(e2wtypes.WalletLocker).Unlock(context.Background(), []byte("pass")))
	unlocked, err = wallet.(e2wtypes.WalletLocker).IsUnlocked(context.Background())
	require.NoError(t, err)
	require.True(t, unlocked)
	// Lock
	require.NoError(t, wallet.(e2wtypes.WalletLocker).Lock(context.Background()))
	unlocked, err = wallet.(e2wtypes.WalletLocker).IsUnlocked(context.Background())
	require.NoError(t, err)
	require.False(t, unlocked)
}