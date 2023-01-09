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
	id, err := getKSUIDString()
	if err != nil {
		return "", fmt.Errorf("Random Service: cannot generate id with error %w", err)
	}

	return id, nil
}

func (s *Service) GenerateRandomValue() (int64, error) {
	n, err := getRandomNumber()
	if err != nil {
		return 0, fmt.Errorf("Random Service: cannot generate random value with error %w", err)
	}

	return n, nil
}

func (s *Service) SaveRandomValue(id string, value int64) error {
	err := s.storage.AddEntry(models.RandomEntity{GenerationID: id, RandomValue: value})
	if err != nil {
		return fmt.Errorf("Random Service: cannot save value with id %s", id)
	}

	return nil
}

func (s *Service) RetrieveRandomValue(id string) (models.RandomEntity, error) {
	re := s.storage.GetEntryByID(id)
	if re.IsEmpty() {
		return re, fmt.Errorf("Random Service: cannot retrieve value with id %s", id)
	}

	return re, nil
}
