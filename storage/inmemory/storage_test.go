package inmemory_test

import (
	"testing"

	"github.com/sbxb/av-random/models"
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
	wantEmpty := true

	entry := s.storage.GetEntryByID(key)
	gotEmpty := entry.IsEmpty()

	s.Equal(wantEmpty, gotEmpty)
}

func (s *MemoryStorageTestSuite) Test_02_MemoryStorage_AddThenGetEntry() {
	key := "some_key"
	wantEntry := models.RandomEntity{
		GenerationID: key,
		RandomValue:  555,
	}

	err := s.storage.AddEntry(wantEntry)
	s.NoError(err)

	gotEntry := s.storage.GetEntryByID(key)
	s.Equal(wantEntry, gotEntry)
}

func TestMemoryStorageTestSuite(t *testing.T) {
	suite.Run(t, new(MemoryStorageTestSuite))
}
