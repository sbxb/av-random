package random

import (
	"context"
	"errors"
	"fmt"

	"github.com/sbxb/av-random/models"
	"github.com/sbxb/av-random/storage"
)

type Service struct {
	storage storage.Storage
	gen     RandomGenerator
}

func New(storage storage.Storage) (*Service, error) {
	return &Service{
		storage: storage,
		gen:     RandomGenerator{},
	}, nil
}

func (s *Service) GenerateID() (string, error) {
	id := s.gen.GenerateKSUID()
	if id == "" {
		return "", fmt.Errorf("Random Service: cannot generate id")
	}

	return id, nil
}

func (s *Service) GenerateRandomValue(valueType string, length int) (string, error) {
	var value string

	switch valueType {
	case "dec":
		value = s.gen.GenerateDec(length)
	case "hex":
		value = s.gen.GenerateHex(length)
	case "str":
		value = s.gen.GenerateStr(length)
	case "stralnum":
		value = s.gen.GenerateStrAlnum(length)
	case "uuid":
		value = s.gen.GenerateUUID()
	}

	if value == "" {
		return "", fmt.Errorf("Random Service: cannot generate random value")
	}

	return value, nil
}

func (s *Service) SaveRandomValue(ctx context.Context, entity models.RandomEntity) error {
	err := s.storage.AddEntry(ctx, entity)
	if err != nil {
		return fmt.Errorf("Random Service: cannot save value %v", entity)
	}

	return nil
}

func (s *Service) RetrieveRandomValue(ctx context.Context, id string) (models.RandomEntity, error) {
	re, err := s.storage.GetEntryByID(ctx, id)
	if err != nil {
		if errors.Is(err, storage.ErrEntryNotFound) {
			return re, fmt.Errorf("Random Service: cannot find value with id %s", id)
		}
		return re, fmt.Errorf("Random Service: cannot retrieve value with id %s due to internal error %w", id, err)
	}

	return re, nil
}
