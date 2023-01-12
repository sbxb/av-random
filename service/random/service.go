package random

import (
	"context"
	"fmt"

	"github.com/sbxb/av-random/models"
	"github.com/sbxb/av-random/storage"
)

type Service struct {
	storage storage.Storage
}

func New(storage storage.Storage) (*Service, error) {
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

func (s *Service) SaveRandomValue(ctx context.Context, id string, value int64) error {
	err := s.storage.AddEntry(ctx, models.RandomEntity{GenerationID: id, RandomValue: value})
	if err != nil {
		return fmt.Errorf("Random Service: cannot save value with id %s", id)
	}

	return nil
}

func (s *Service) RetrieveRandomValue(ctx context.Context, id string) (models.RandomEntity, error) {
	re, err := s.storage.GetEntryByID(ctx, id)
	if err != nil {
		return re, fmt.Errorf("Random Service: cannot retrieve value with id %s due to internal error %w", id, err)
	}
	if re.IsEmpty() {
		return re, fmt.Errorf("Random Service: cannot find value with id %s", id)
	}

	return re, nil
}
