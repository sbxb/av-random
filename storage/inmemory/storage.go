package inmemory

import (
	"context"
	"sync"

	"github.com/sbxb/av-random/models"
	"github.com/sbxb/av-random/storage"
)

type MemoryStorage struct {
	sync.RWMutex

	data map[string]models.RandomEntity
}

func NewMemoryStorage() (*MemoryStorage, error) {
	data := make(map[string]models.RandomEntity)

	return &MemoryStorage{data: data}, nil // Never returns an error, need it for compatibility
}

func (ms *MemoryStorage) AddEntry(ctx context.Context, entry models.RandomEntity) error {
	ms.Lock()
	defer ms.Unlock()

	ms.data[entry.GenerationID] = entry

	return nil // Never returns an error, need it for compatibility
}

func (ms *MemoryStorage) GetEntryByID(ctx context.Context, id string) (models.RandomEntity, error) {
	ms.RLock()
	defer ms.RUnlock()

	res, ok := ms.data[id]
	if !ok {
		return res, storage.ErrEntryNotFound
	}

	return res, nil
}
