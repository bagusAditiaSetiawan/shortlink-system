package redis

import "time"

type RedisService interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (interface{}, error)
}
