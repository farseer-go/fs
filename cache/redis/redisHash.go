package redis

type redisHash struct {
}

// Set 添加
func (redisHash *redisHash) Set(key string, values ...interface{}) error {
	return rdb.HSet(ctx, key, values).Err()
}

// SetMap 添加Map
func (redisHash *redisHash) SetMap(key string, values map[string]string) error {
	for k, v := range values {
		err := rdb.HSet(ctx, key, k, v).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

// Get 获取
func (redisHash *redisHash) Get(key string, fields string) (string, error) {
	return rdb.HGet(ctx, key, fields).Result()
}

// GetAll 获取所有集合数据
func (redisHash *redisHash) GetAll(key string) (map[string]string, error) {
	return rdb.HGetAll(ctx, key).Result()
}

// Exists 成员是否存在
func (redisHash *redisHash) Exists(key string, field string) (bool, error) {
	return rdb.HExists(ctx, key, field).Result()
}

// Remove 移除指定成员
func (redisHash *redisHash) Remove(key string, field string) (bool, error) {
	result, err := rdb.HDel(ctx, key, field).Result()
	return result > 0, err
}
