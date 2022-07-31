package redis

import "time"

type redisString struct {
}

// Set 设置缓存
func (redisString *redisString) Set(key string, value string) error {
	return rdb.Set(ctx, key, value, 0).Err()
}

// Get 获取缓存
func (redisString *redisString) Get(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

// SetNX 设置过期时间
func (redisString *redisString) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	return rdb.SetNX(ctx, key, value, expiration).Result()
}

// TTL 获取过期时间
func (redisString *redisString) TTL(key string) (time.Duration, error) {
	return rdb.TTL(ctx, key).Result()
}
