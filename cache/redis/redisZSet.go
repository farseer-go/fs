package redis

import (
	"github.com/go-redis/redis/v8"
)

type redisZSet struct {
	rdb *redis.Client
}

type redisZ struct {
	Score  float64
	Member interface{}
}
type redisZRangeBy struct {
	Min, Max      string
	Offset, Count int64
}

// Add 添加
func (redisZSet *redisZSet) Add(key string, members ...*redisZ) (bool, error) {
	var redisZZ []*redis.Z
	for _, member := range members {
		redisZZ = append(redisZZ, &redis.Z{Score: member.Score, Member: member.Member})
	}
	result, err := redisZSet.rdb.ZAdd(ctx, key, redisZZ...).Result()
	return result > 0, err
}

// Score 获取指定成员score
func (redisZSet *redisZSet) Score(key string, member string) (float64, error) {
	return redisZSet.rdb.ZScore(ctx, key, member).Result()
}

// Range 获取有序集合指定区间内的成员
func (redisZSet *redisZSet) Range(key string, start int64, stop int64) ([]string, error) {
	return redisZSet.rdb.ZRange(ctx, key, start, stop).Result()
}

// RevRange 获取有序集合指定区间内的成员分数从高到低
func (redisZSet *redisZSet) RevRange(key string, start int64, stop int64) ([]string, error) {
	return redisZSet.rdb.ZRevRange(ctx, key, start, stop).Result()
}

// RangeByScore 获取指定分数区间的成员列表
func (redisZSet *redisZSet) RangeByScore(key string, opt *redisZRangeBy) ([]string, error) {
	rby := redis.ZRangeBy{Min: opt.Min, Max: opt.Max, Offset: opt.Offset, Count: opt.Count}
	return redisZSet.rdb.ZRangeByScore(ctx, key, &rby).Result()
}
