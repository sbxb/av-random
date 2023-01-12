package storage

import (
	"context"

	"github.com/sbxb/av-random/models"
)

type Storage interface {
	AddEntry(ctx context.Context, entry models.RandomEntity) error
	GetEntryByID(ctx context.Context, id string) (models.RandomEntity, error)
}
