package indexer_test

import (
	"sync"
	"testing"

	"github.com/alex-necsoiu/go-eth-wallet/indexer"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIndex(t *testing.T) {
	index := indexer.New()

	id := uuid.New()
	name := "Test name"

	assert.False(t, index.IDKnown(id))
	assert.False(t, index.NameKnown(name))
	index.Add(id, name)
	assert.True(t, index.IDKnown(id))
	assert.True(t, index.NameKnown(name))

	foundID, exists := index.ID(name)
	assert.True(t, exists)
	assert.Equal(t, id, foundID)

	foundName, exists := index.Name(id)
	assert.True(t, exists)
	assert.Equal(t, name, foundName)

	index.Remove(id, name)

	assert.False(t, index.IDKnown(id))
	assert.False(t, index.NameKnown(name))
}

func TestIndexSerDeser(t *testing.T) {
	index := indexer.New()

	id1 := uuid.New()
	name1 := "Test name 1"
	index.Add(id1, name1)
	id2 := uuid.New()
	name2 := "Test name 2"
	index.Add(id2, name2)

	ser, err := index.Serialize()
	require.Nil(t, err)

	index, err = indexer.Deserialize(ser)
	require.Nil(t, err)

	foundName, exists := index.Name(id1)
	assert.True(t, exists)
	assert.Equal(t, foundName, name1)
	foundID, exists := index.ID(name1)
	assert.True(t, exists)
	assert.Equal(t, foundID, id1)

	foundName, exists = index.Name(id2)
	assert.True(t, exists)
	assert.Equal(t, foundName, name2)
	foundID, exists = index.ID(name2)
	assert.True(t, exists)
	assert.Equal(t, foundID, id2)
}

func TestConcurrency(t *testing.T) {
	index := indexer.New()

	// Create a number of runners that will try to add and remove indices simultaneously.
	var runWG sync.WaitGroup
	var setupWG sync.WaitGroup
	starter := make(chan any)
	for i := 0; i < 64; i++ {
		setupWG.Add(1)
		runWG.Add(1)
		go func() {
			setupWG.Done()
			id, err := uuid.NewRandom()
			require.NoError(t, err)
			name := id.String()
			<-starter
			index.Add(id, name)
			require.True(t, index.IDKnown(id))
			require.True(t, index.NameKnown(name))
			index.Remove(id, name)
			runWG.Done()
		}()
	}
	// Wait for setup to complete.
	setupWG.Wait()
	// Start the jobs by closing the channel.
	close(starter)

	// Wait for run to complete
	runWG.Wait()
}
