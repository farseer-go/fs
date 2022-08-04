package redis

import (
	"github.com/go-redis/redis/v8"
	"time"
)

// 分布式锁
type redisLock struct {
	rdb *redis.Client
}

type lockResult struct {
	key        string
	expiration time.Duration
	rdb        *redis.Client
}

// GetLocker 获得一个锁
func (r redisLock) GetLocker(key string, expiration time.Duration) lockResult {
	return lockResult{
		rdb:        r.rdb,
		key:        key,
		expiration: expiration,
	}
}

// TryLock 尝试加锁
func (r *lockResult) TryLock() bool {
	cmd := r.rdb.SetNX(ctx, r.key, 1, r.expiration)
	result, _ := cmd.Result()
	return result
}

// ReleaseLock 锁放锁
func (r *lockResult) ReleaseLock() {
	r.rdb.Del(ctx, r.key)
}
