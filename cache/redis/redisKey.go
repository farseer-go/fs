package redis

import (
	"github.com/go-redis/redis/v8"
	"time"
)

type redisKey struct {
	rdb *redis.Client
}

// TTL 获取过期时间
func (redisKey *redisKey) TTL(key string) (time.Duration, error) {
	return redisKey.rdb.TTL(ctx, key).Result()
}

// Del 删除
func (redisKey *redisKey) Del(keys ...string) (bool, error) {
	result, err := redisKey.rdb.Del(ctx, keys...).Result()
	return result > 0, err
}

// Exists key值是否存在
func (redisKey *redisKey) Exists(keys ...string) (bool, error) {
	result, err := redisKey.rdb.Exists(ctx, keys...).Result()
	return result > 0, err
}
