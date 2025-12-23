package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, ttl time.Duration) error
}

type redisCache struct{
	client *redis.Client
}

func NewRedisCache(redisAddr string) Cache {
	if redisAddr == "" {
		// return noop cache
		return &redisCache{client: redis.NewClient(&redis.Options{Addr: "127.0.0.1:0"})}
	}
	opt := &redis.Options{Addr: redisAddr}
	client := redis.NewClient(opt)
	return &redisCache{client: client}
}

func (r *redisCache) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *redisCache) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}
