package random

import (
	"fmt"

	"github.com/sbxb/av-random/models"
)

type Storage interface {
	AddEntry(entry models.RandomEntity) error
	GetEntryByID(id string) models.RandomEntity
}

type Service struct {
	storage Storage
}

func New(storage Storage) (*Service, error) {
	return &Service{
		storage: storage,
	}, nil
}

func (s *Service) GenerateID() (string, error) {
	id := getKSUIDString()
	if id == "" {
		return id, fmt.Errorf("Random Service: cannot generate id")
	}

	return id, nil
}

func (s *Service) GenerateRandomValue() (int64, error) {
	return getRandomNumber(), nil
}
