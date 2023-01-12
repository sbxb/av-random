package random_test

import (
	"testing"

	"github.com/sbxb/av-random/service/random"
	"github.com/stretchr/testify/suite"
)

type RandomGeneratorTestSuite struct {
	suite.Suite
	gen random.RandomGenerator
}

func (s *RandomGeneratorTestSuite) SetupSuite() {
	s.gen = random.RandomGenerator{}
}

func (s *RandomGeneratorTestSuite) TearDownSuite() {
	// empty
}

func (s *RandomGeneratorTestSuite) Test_01_RandomGenerator_Digits() {
	for i := 0; i < 20; i++ {
		s.T().Log(s.gen.GenerateDec(i))
	}
}

func (s *RandomGeneratorTestSuite) Test_02_RandomGenerator_HexDigits() {
	for i := 0; i < 20; i++ {
		s.T().Log(s.gen.GenerateHex(i))
	}
}

func (s *RandomGeneratorTestSuite) Test_03_RandomGenerator_String() {
	for i := 0; i < 20; i++ {
		s.T().Log(s.gen.GenerateStr(i))
	}
}

func (s *RandomGeneratorTestSuite) Test_04_RandomGenerator_AlnumString() {
	for i := 0; i < 20; i++ {
		s.T().Log(s.gen.GenerateStrAlnum(i))
	}
}

func (s *RandomGeneratorTestSuite) Test_05_RandomGenerator_UUID() {
	for i := 0; i < 20; i++ {
		s.T().Log(s.gen.GenerateUUID())
	}
}

func (s *RandomGeneratorTestSuite) Test_06_RandomGenerator_KSUID() {
	wantLength := 27 // by design of KSUID

	id := s.gen.GenerateKSUID()

	s.Equal(len(id), wantLength)
}

func TestRandomGeneratorTestSuite(t *testing.T) {
	suite.Run(t, new(RandomGeneratorTestSuite))
}
