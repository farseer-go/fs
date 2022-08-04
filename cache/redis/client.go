package redis

import (
	"context"
	"github.com/farseernet/farseer.go/configure"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	Key    *redisKey
	String *redisString
	Hash   *redisHash
	List   *redisList
	Set    *redisSet
	ZSet   *redisZSet
	Lock   *redisLock
}

// 上下文定义
var ctx = context.Background()

// NewClient 初始化
func NewClient(redisName string) *Client {
	configString := configure.GetString("Redis." + redisName)
	if configString == "" {
		panic("[farseer.yaml]找不到相应的配置：Redis." + redisName)
	}
	redisConfig := configure.ParseConfig[redisConfig](configString)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Server,   //localhost:6379
		Password: redisConfig.Password, // no password Set
		DB:       redisConfig.DB,       // use default DB
	})
	key := &redisKey{rdb: rdb}
	str := &redisString{rdb: rdb}
	hash := &redisHash{rdb: rdb}
	list := &redisList{rdb: rdb}
	set := &redisSet{rdb: rdb}
	zset := &redisZSet{rdb: rdb}
	lock := &redisLock{rdb: rdb}
	return &Client{Key: key, String: str, Hash: hash, List: list, Set: set, ZSet: zset, Lock: lock}
}
