package redis

import (
	"context"
	"encoding/json"
	"fmt"

	r "github.com/go-redis/redis/v9"
	"github.com/sbxb/av-random/config"
	"github.com/sbxb/av-random/models"
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
		return nil, fmt.Errorf("Redis storage: can not connect to DB: %w", err)
	}

	return &RedisStorage{client: cl}, nil
}

// FIXME Errors should be of custom types (placed in storage/errors.go) !!!

func (rs *RedisStorage) AddEntry(ctx context.Context, entry models.RandomEntity) error {
	jEntry, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	err = rs.client.Set(ctx, entry.GenerationID, jEntry, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func (rs *RedisStorage) GetEntryByID(ctx context.Context, id string) (models.RandomEntity, error) {
	var res models.RandomEntity

	str, err := rs.client.Get(ctx, id).Result()
	if err != nil {
		if err == r.Nil {
			return res, nil
		}
		return res, fmt.Errorf("Redis storage: can not extract entry: %w", err)
	}

	err = json.Unmarshal([]byte(str), &res)
	if err != nil {
		return res, fmt.Errorf("Redis storage: can not unmarshal extracted entry: %w", err)
	}

	return res, nil
}
