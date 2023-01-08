package inmemory

import (
	"sync"

	"github.com/sbxb/av-random/models"
)

type MemoryStorage struct {
	sync.RWMutex

	data map[string]models.RandomEntity
}

func NewMemoryStorage() (*MemoryStorage, error) {
	data := make(map[string]models.RandomEntity)

	// Never returns an error, need it for compatibility with other possible storages
	return &MemoryStorage{data: data}, nil
}

func (ms *MemoryStorage) AddEntry(entry models.RandomEntity) error {
	ms.Lock()
	defer ms.Unlock()

	ms.data[entry.GenerationID] = entry

	return nil
}

func (ms *MemoryStorage) GetEntryByID(id string) models.RandomEntity {
	ms.Lock()
	defer ms.Unlock()

	return ms.data[id]
}
