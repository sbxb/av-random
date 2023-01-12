package random_test

import (
	"context"
	"testing"

	"github.com/sbxb/av-random/service/random"
	"github.com/sbxb/av-random/storage/inmemory"
	"github.com/stretchr/testify/suite"
)

type RandomServiceTestSuite struct {
	suite.Suite
	service *random.Service
}

func (s *RandomServiceTestSuite) SetupSuite() {
	storage, _ := inmemory.NewMemoryStorage()
	service, _ := random.New(storage)

	s.service = service
}

func (s *RandomServiceTestSuite) TearDownSuite() {
	// empty
}

func (s *RandomServiceTestSuite) Test_01_RandomService_SaveRandomValue() {
	id := "someid"
	value := int64(555)

	err := s.service.SaveRandomValue(context.Background(), id, value)
	s.NoError(err)
}

func (s *RandomServiceTestSuite) Test_02_RandomService_RetrieveRandomValue() {
	id := "someid"
	wantValue := int64(555)

	got, err := s.service.RetrieveRandomValue(context.Background(), id)
	s.NoError(err)

	s.Equal(wantValue, got.RandomValue)
}

func TestRandomServiceTestSuite(t *testing.T) {
	suite.Run(t, new(RandomServiceTestSuite))
}
