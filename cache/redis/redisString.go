package redis

import (
	"github.com/go-redis/redis/v8"
	"time"
)

type redisString struct {
	rdb *redis.Client
}

// Set 设置缓存
func (redisString *redisString) Set(key string, value interface{}) error {
	return redisString.rdb.Set(ctx, key, value, 0).Err()
}

// Get 获取缓存
func (redisString *redisString) Get(key string) (string, error) {
	return redisString.rdb.Get(ctx, key).Result()
}

// SetNX 设置过期时间
func (redisString *redisString) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return redisString.rdb.SetNX(ctx, key, value, expiration).Result()
}

// TTL 获取过期时间
func (redisString *redisString) TTL(key string) (time.Duration, error) {
	return redisString.rdb.TTL(ctx, key).Result()
}

// Remove 删除
func (redisString *redisString) Remove(keys ...string) (bool, error) {
	result, err := redisString.rdb.Del(ctx, keys...).Result()
	return result > 0, err
}

// Exists key值是否存在
func (redisString *redisString) Exists(keys ...string) (bool, error) {
	result, err := redisString.rdb.Exists(ctx, keys...).Result()
	return result > 0, err
}
