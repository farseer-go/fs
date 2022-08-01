package redis

import "github.com/go-redis/redis/v8"

type redisHash struct {
	rdb *redis.Client
}

// Set 添加
func (redisHash *redisHash) Set(key string, values ...interface{}) error {
	return redisHash.rdb.HSet(ctx, key, values).Err()
}

// SetMap 添加Map
func (redisHash *redisHash) SetMap(key string, values map[string]string) error {
	for k, v := range values {
		err := redisHash.rdb.HSet(ctx, key, k, v).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

// Get 获取
func (redisHash *redisHash) Get(key string, field string) (string, error) {
	return redisHash.rdb.HGet(ctx, key, field).Result()
}

// GetAll 获取所有集合数据
func (redisHash *redisHash) GetAll(key string) (map[string]string, error) {
	return redisHash.rdb.HGetAll(ctx, key).Result()
}

// Exists 成员是否存在
func (redisHash *redisHash) Exists(key string, field string) (bool, error) {
	return redisHash.rdb.HExists(ctx, key, field).Result()
}

// Del 移除指定成员
func (redisHash *redisHash) Del(key string, fields ...string) (bool, error) {
	result, err := redisHash.rdb.HDel(ctx, key, fields...).Result()
	return result > 0, err
}
