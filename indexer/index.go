// Package indexer provides an indexing system between names and UUIDs.
package indexer

import (
	"encoding/json"
	"sync"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Index maps between names and IDs.
type Index struct {
	names map[uuid.UUID]string
	ids   map[string]uuid.UUID
	mutex sync.RWMutex
}

// New creates a new index.
func New() *Index {
	return &Index{
		names: make(map[uuid.UUID]string),
		ids:   make(map[string]uuid.UUID),
		mutex: sync.RWMutex{},
	}
}

// Add adds an entry to this index.
func (i *Index) Add(id uuid.UUID, name string) {
	i.mutex.Lock()
	i.names[id] = name
	i.ids[name] = id
	i.mutex.Unlock()
}

// Remove removes an entry from this index.
func (i *Index) Remove(id uuid.UUID, name string) {
	i.mutex.Lock()
	delete(i.names, id)
	delete(i.ids, name)
	i.mutex.Unlock()
}

// Name fetches the name of an entry given its ID.
// If present the second return value will be true, otherwise false.
func (i *Index) Name(id uuid.UUID) (string, bool) {
	i.mutex.RLock()
	res, exists := i.names[id]
	i.mutex.RUnlock()

	return res, exists
}

// NameKnown returns true if this name is known.
func (i *Index) NameKnown(name string) bool {
	i.mutex.RLock()
	_, exists := i.ids[name]
	i.mutex.RUnlock()

	return exists
}

// ID fetches the ID of an entry given its name.
// If present the second return value will be true, otherwise false.
func (i *Index) ID(name string) (uuid.UUID, bool) {
	i.mutex.RLock()
	res, exists := i.ids[name]
	i.mutex.RUnlock()

	return res, exists
}

// IDKnown returns true if this ID is known.
func (i *Index) IDKnown(id uuid.UUID) bool {
	i.mutex.RLock()
	_, exists := i.names[id]
	i.mutex.RUnlock()

	return exists
}

type indexEntry struct {
	ID   uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

// Deserialize deserializes a serialized index.
func Deserialize(data []byte) (*Index, error) {
	var entries []*indexEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal data")
	}

	index := New()
	for _, entry := range entries {
		index.Add(entry.ID, entry.Name)
	}

	return index, nil
}

// Serialize serializes an index.
func (i *Index) Serialize() ([]byte, error) {
	entries := make([]*indexEntry, 0)

	i.mutex.RLock()
	for k, v := range i.names {
		entries = append(entries, &indexEntry{
			ID:   k,
			Name: v,
		})
	}
	i.mutex.RUnlock()

	data, err := json.Marshal(entries)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal data")
	}

	return data, nil
}
