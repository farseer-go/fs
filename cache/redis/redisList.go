package redis

import (
	"github.com/go-redis/redis/v8"
	"time"
)

type redisList struct {
	rdb *redis.Client
}

// Push 添加
func (redisList *redisList) Push(key string, values ...interface{}) (bool, error) {
	result, err := redisList.rdb.RPush(ctx, key, values).Result()
	return result > 0, err
}

// Set 设置指定索引的值
func (redisList *redisList) Set(key string, index int64, value interface{}) (string, error) {
	return redisList.rdb.LSet(ctx, key, index, value).Result()
}

// Rem 移除指定数量的value，count=0 移除全部，其他移除指定数量的
func (redisList *redisList) Rem(key string, count int64, value interface{}) (bool, error) {
	result, err := redisList.rdb.LRem(ctx, key, count, value).Result()
	return result > 0, err

}

// Len 获取长度
func (redisList *redisList) Len(key string) (int64, error) {
	return redisList.rdb.LLen(ctx, key).Result()
}

// Range 遍历
func (redisList *redisList) Range(key string, start int64, stop int64) ([]string, error) {
	return redisList.rdb.LRange(ctx, key, start, stop).Result()
}

// BLPop 命令移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
func (redisList *redisList) BLPop(timeout time.Duration, keys ...string) ([]string, error) {
	return redisList.rdb.BLPop(ctx, timeout, keys...).Result()
}
