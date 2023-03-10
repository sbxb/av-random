package inmemory_test

import (
	"context"
	"testing"

	"github.com/sbxb/av-random/models"
	"github.com/sbxb/av-random/storage"
	"github.com/sbxb/av-random/storage/inmemory"
	"github.com/stretchr/testify/suite"
)

type MemoryStorageTestSuite struct {
	suite.Suite
	storage *inmemory.MemoryStorage
}

func (s *MemoryStorageTestSuite) SetupSuite() {
	s.storage, _ = inmemory.NewMemoryStorage()
}

func (s *MemoryStorageTestSuite) TearDownSuite() {
	// empty
}

func (s *MemoryStorageTestSuite) Test_01_MemoryStorage_GetNonexistentEntry() {
	key := "nonexistent"

	_, err := s.storage.GetEntryByID(context.Background(), key)

	s.ErrorIs(err, storage.ErrEntryNotFound)
}

func (s *MemoryStorageTestSuite) Test_02_MemoryStorage_AddThenGetEntry() {
	key := "some_key"
	wantEntry := models.RandomEntity{
		GenerationID:    key,
		RandomValue:     "555",
		RandomValueType: "dec",
	}
	ctx := context.Background()

	err := s.storage.AddEntry(ctx, wantEntry)
	s.NoError(err)

	gotEntry, err := s.storage.GetEntryByID(ctx, key)
	s.NoError(err)

	s.Equal(wantEntry, gotEntry)
}

func TestMemoryStorageTestSuite(t *testing.T) {
	suite.Run(t, new(MemoryStorageTestSuite))
}
