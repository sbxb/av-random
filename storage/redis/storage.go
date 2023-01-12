package redis

import (
	"context"
	"encoding/json"
	"fmt"

	r "github.com/go-redis/redis/v9"
	"github.com/sbxb/av-random/config"
	"github.com/sbxb/av-random/models"
	"github.com/sbxb/av-random/storage"
)

type RedisStorage struct {
	client *r.Client
}

func NewRedisStorage(cfg config.Redis) (*RedisStorage, error) {
	cl := r.NewClient(&r.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       0,
	})

	_, err := cl.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("Redis Storage: can not connect to DB: %w", err)
	}

	return &RedisStorage{client: cl}, nil
}

func (rs *RedisStorage) AddEntry(ctx context.Context, entry models.RandomEntity) error {
	jEntry, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("Redis Storage: AddEntry failed: %w", err)
	}

	err = rs.client.Set(ctx, entry.GenerationID, jEntry, 0).Err()
	if err != nil {
		return fmt.Errorf("Redis Storage: AddEntry failed: %w", err)
	}

	return nil
}

func (rs *RedisStorage) GetEntryByID(ctx context.Context, id string) (models.RandomEntity, error) {
	var res models.RandomEntity

	str, err := rs.client.Get(ctx, id).Result()
	if err != nil {
		if err == r.Nil {
			return res, storage.ErrEntryNotFound
		}
		return res, fmt.Errorf("Redis Storage: GetEntryByID failed: %w", err)
	}

	err = json.Unmarshal([]byte(str), &res)
	if err != nil {
		return res, fmt.Errorf("Redis Storage: GetEntryByID failed: %w", err)
	}

	return res, nil
}
