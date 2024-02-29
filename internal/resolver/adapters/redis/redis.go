package redis

import (
	"context"
	"github.com/red-life/zone/internal/resolver"
	"github.com/redis/go-redis/v9"
	"time"
)

var _ resolver.Cache = (*RedisCache)(nil)

func NewRedisCache(rdb *redis.Client, ttl time.Duration) *RedisCache {
	return &RedisCache{
		rdb: rdb,
		ttl: ttl,
	}
}

type RedisCache struct {
	rdb *redis.Client
	ttl time.Duration
}

func (r *RedisCache) Set(key string, value string) error {
	err := r.rdb.SetEx(context.Background(), key, value, r.ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisCache) Get(key string) (string, error) {
	var val *redis.StringCmd
	_, err := r.rdb.Pipelined(context.Background(), func(pipe redis.Pipeliner) error {
		value := pipe.Get(context.Background(), key)
		if value.Err() != nil {
			return value.Err()
		}
		val = value
		if err := pipe.Expire(context.Background(), key, r.ttl).Err(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return val.Val(), nil
}

func (r *RedisCache) Delete(key string) error {
	return r.rdb.Del(context.Background(), key).Err()
}
