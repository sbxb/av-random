package random

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/segmentio/ksuid"
)

const (
	MaxLength = 4096
)

const (
	Digits    = "0123456789"
	HexDigits = "0123456789abcdef"
	Letters   = "abcdefghijklmnopqrstuvwxyz"
	Alnum     = Digits + Letters
)

type RandomGenerator struct {
}

func generateSequence(length int, source string) string {
	if length < 1 {
		length = 1
	}

	if length > MaxLength {
		length = MaxLength
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = source[rand.Intn(len(source))]
	}

	return string(b)
}

func (rg *RandomGenerator) GenerateDec(length int) string {
	return generateSequence(length, Digits)
}

func (rg *RandomGenerator) GenerateHex(length int) string {
	return generateSequence(length, HexDigits)
}

func (rg *RandomGenerator) GenerateStr(length int) string {
	return generateSequence(length, Letters)
}

func (rg *RandomGenerator) GenerateStrAlnum(length int) string {
	return generateSequence(length, Alnum)
}

func (rg *RandomGenerator) GenerateUUID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		return ""
	}

	return id.String()
}

func (rg *RandomGenerator) GenerateKSUID() string {
	id, err := ksuid.NewRandom()
	if err != nil {
		return ""
	}

	return id.String()
}
