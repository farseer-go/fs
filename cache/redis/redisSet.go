package redis

import "github.com/go-redis/redis/v8"

type redisSet struct {
	rdb *redis.Client
}

// Add 添加
func (redisSet *redisSet) Add(key string, members ...interface{}) (bool, error) {
	result, err := redisSet.rdb.SAdd(ctx, key, members...).Result()
	return result > 0, err
}

// Card 获取数量
func (redisSet *redisSet) Card(key string) (int64, error) {
	return redisSet.rdb.SCard(ctx, key).Result()
}

// Rem 移除指定成员
func (redisSet *redisSet) Rem(key string, members ...interface{}) (bool, error) {
	result, err := redisSet.rdb.SRem(ctx, key, members...).Result()
	return result > 0, err
}

// Members 获取所有成员
func (redisSet *redisSet) Members(key string) ([]string, error) {
	return redisSet.rdb.SMembers(ctx, key).Result()
}

// IsMember 判断指定成员是否存在
func (redisSet *redisSet) IsMember(key string, member interface{}) (bool, error) {
	return redisSet.rdb.SIsMember(ctx, key, member).Result()
}

// Diff 获取差集
func (redisSet *redisSet) Diff(keys ...string) ([]string, error) {
	return redisSet.rdb.SDiff(ctx, keys...).Result()
}

// DiffStore 将差集，保存在指定集合中
func (redisSet *redisSet) DiffStore(destination string, keys ...string) (bool, error) {
	result, err := redisSet.rdb.SDiffStore(ctx, destination, keys...).Result()
	return result > 0, err
}

// Inter 获取交集
func (redisSet *redisSet) Inter(keys ...string) ([]string, error) {
	return redisSet.rdb.SInter(ctx, keys...).Result()
}

// InterStore 将交集，保存在指定集合中
func (redisSet *redisSet) InterStore(destination string, keys ...string) (bool, error) {
	result, err := redisSet.rdb.SInterStore(ctx, destination, keys...).Result()
	return result > 0, err
}

// Union 获取并集
func (redisSet *redisSet) Union(keys ...string) ([]string, error) {
	return redisSet.rdb.SUnion(ctx, keys...).Result()
}

// UnionStore 将并集，保存在指定集合中
func (redisSet *redisSet) UnionStore(destination string, keys ...string) (bool, error) {
	result, err := redisSet.rdb.SUnionStore(ctx, destination, keys...).Result()
	return result > 0, err
}
