package cache

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisCacheConnection struct {
	client *redis.Client
	ttl    time.Duration
}

func NewRedisCacheConnection(redisAddr string, ttl time.Duration) Provider {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})
	return &redisCacheConnection{
		client: rdb,
		ttl:    ttl,
	}
}

func (r *redisCacheConnection) Set(ctx context.Context, key string, value []byte) error {
	return r.client.Set(ctx, key, value, r.ttl).Err()
}

func (r *redisCacheConnection) Get(ctx context.Context, key string) ([]byte, error) {
	raw, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, ErrCacheMiss
		}
		return nil, err
	}
	return raw, nil
}

func (r *redisCacheConnection) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
