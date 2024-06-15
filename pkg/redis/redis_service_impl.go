package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

type RedisServiceImpl struct {
	RedisClient *redis.Client
}

func NewRedisServiceImpl(addr string, password string, db int) *RedisServiceImpl {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,     // Redis server address
		Password: password, // No password set
		DB:       db,       // Use default DB
	})
	return &RedisServiceImpl{
		RedisClient: client,
	}
}

func (service *RedisServiceImpl) Set(key string, value interface{}, expiration time.Duration) error {
	err := service.RedisClient.Set(ctx, key, value, expiration).Err()
	return err
}

func (service *RedisServiceImpl) Get(key string) (interface{}, error) {
	result, err := service.RedisClient.Get(ctx, key).Result()
	return result, err
}
