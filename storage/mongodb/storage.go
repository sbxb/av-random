package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/sbxb/av-random/config"
	"github.com/sbxb/av-random/models"
	"github.com/sbxb/av-random/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoStorage struct {
	client  *mongo.Client
	entries *mongo.Collection
}

func NewMongoStorage(cfg config.MongoDB) (*MongoStorage, error) {
	ms := &MongoStorage{}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s", cfg.User, cfg.Password, cfg.Address, cfg.DBName)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("MongoDB Storage: can not connect to server: %w", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("MongoDB Storage: can not ping the server: %w", err)
	}

	ms.client = client
	ms.entries = client.Database(cfg.DBName).Collection("entries")

	return ms, nil
}

func (ms *MongoStorage) AddEntry(ctx context.Context, entry models.RandomEntity) error {
	newEntry := bson.D{
		{Key: "GenerationID", Value: entry.GenerationID},
		{Key: "RandomValue", Value: entry.RandomValue},
		{Key: "RandomValueType", Value: entry.RandomValueType},
	}

	_, err := ms.entries.InsertOne(ctx, newEntry)
	if err != nil {
		return fmt.Errorf("MongoDB Storage: AddEntry failed: %w", err)
	}

	return nil
}

func (ms *MongoStorage) GetEntryByID(ctx context.Context, id string) (models.RandomEntity, error) {
	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{
					{Key: "GenerationID", Value: bson.D{{Key: "$eq", Value: id}}},
				},
			},
		},
	}

	var res models.RandomEntity

	if err := ms.entries.FindOne(ctx, filter).Decode(&res); err != nil {
		if err == mongo.ErrNoDocuments {
			return res, storage.ErrEntryNotFound
		}
		return res, fmt.Errorf("MongoDB Storage: GetEntryByID failed: %w", err)
	}

	return res, nil
}
