package redis_test

import (
	"testing"

	"github.com/sbxb/av-random/models"
	"github.com/sbxb/av-random/storage/redis"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RedisStorageTestSuite struct {
	suite.Suite
	storage *redis.RedisStorage
}

func (s *RedisStorageTestSuite) SetupSuite() {
	var err error
	s.storage, err = redis.NewRedisStorage()
	require.Nil(s.T(), err)
}

func (s *RedisStorageTestSuite) TearDownSuite() {
	// empty
}

func (s *RedisStorageTestSuite) Test_01_RedisStorage_GetNonexistentEntry() {
	key := "nonexistent"
	wantEmpty := true

	entry, _ := s.storage.GetEntryByID(key)
	gotEmpty := entry.IsEmpty()

	s.Equal(wantEmpty, gotEmpty)
}

func (s *RedisStorageTestSuite) Test_02_RedisStorage_AddEntry() {
	key := "some_key"
	wantEntry := models.RandomEntity{
		GenerationID: key,
		RandomValue:  555,
	}

	err := s.storage.AddEntry(wantEntry)
	s.NoError(err)
}

func (s *RedisStorageTestSuite) Test_03_RedisStorage_GetEntry() {
	key := "some_key"
	wantEntry := models.RandomEntity{
		GenerationID: key,
		RandomValue:  555,
	}

	gotEntry, err := s.storage.GetEntryByID(key)
	s.NoError(err)
	s.Equal(wantEntry, gotEntry)
}

func TestRedisStorageTestSuite(t *testing.T) {
	suite.Run(t, new(RedisStorageTestSuite))
}
