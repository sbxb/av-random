package redis_test

import (
	"context"
	"testing"

	"github.com/sbxb/av-random/config"
	"github.com/sbxb/av-random/models"
	"github.com/sbxb/av-random/storage"
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
	s.storage, err = redis.NewRedisStorage(config.Redis{
		Address:  "localhost:6379",
		Password: "password",
	})
	require.Nil(s.T(), err)
}

func (s *RedisStorageTestSuite) TearDownSuite() {
	// empty
}

func (s *RedisStorageTestSuite) Test_01_RedisStorage_GetNonexistentEntry() {
	key := "nonexistent"
	wantEmpty := true

	entry, err := s.storage.GetEntryByID(context.Background(), key)

	s.ErrorIs(err, storage.ErrEntryNotFound)

	gotEmpty := entry.IsEmpty()
	s.Equal(wantEmpty, gotEmpty)
}

func (s *RedisStorageTestSuite) Test_02_RedisStorage_AddEntry() {
	key := "some_key"
	wantEntry := models.RandomEntity{
		GenerationID: key,
		RandomValue:  555,
	}

	err := s.storage.AddEntry(context.Background(), wantEntry)
	s.NoError(err)
}

func (s *RedisStorageTestSuite) Test_03_RedisStorage_GetEntry() {
	key := "some_key"
	wantEntry := models.RandomEntity{
		GenerationID: key,
		RandomValue:  555,
	}

	gotEntry, err := s.storage.GetEntryByID(context.Background(), key)
	s.NoError(err)
	s.Equal(wantEntry, gotEntry)
}

func TestRedisStorageTestSuite(t *testing.T) {
	suite.Run(t, new(RedisStorageTestSuite))
}
