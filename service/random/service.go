package random

import "github.com/sbxb/av-random/models"

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
