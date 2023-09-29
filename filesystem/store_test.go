package filesystem_test

import (
	"testing"

	filesystem "github.com/alex-necsoiu/go-eth-wallet/filesystem"
	wtypes "github.com/alex-necsoiu/go-eth-wallet/wallet/types"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	store := filesystem.New()
	assert.Equal(t, "filesystem", store.Name())
	store = filesystem.New(filesystem.WithLocation("test"))
	assert.Equal(t, "filesystem", store.Name())
	store = filesystem.New(filesystem.WithLocation("test"), filesystem.WithPassphrase([]byte("secret")))
	assert.Equal(t, "filesystem", store.Name())

	storeLocationProvider, ok := store.(wtypes.StoreLocationProvider)
	assert.True(t, ok)
	assert.Equal(t, "test", storeLocationProvider.Location())
}
