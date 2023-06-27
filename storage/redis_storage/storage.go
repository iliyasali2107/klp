package redis_storage

import (
	"context"
	"fmt"

	"klp/models"

	"github.com/go-redis/redis"
)

type RedisLineStorage struct {
	client *redis.Client
}

func NewRedisLineStorage(client *redis.Client) *RedisLineStorage {
	return &RedisLineStorage{
		client: client,
	}
}

func InitRedisDB(env *models.Env) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	},
	)

	err := rdb.Ping().Err()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to Redis: %w", err)
	}

	return rdb, nil
}

func (kls *RedisLineStorage) Add(ctx context.Context, sport string, cf float64) error {
	return kls.client.Set(sport, cf, 0).Err()
}
