package redis

import (
	"context"
	"github.com/farseernet/farseer.go/configure"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	string *redisString
	hash   *redisHash
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
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DB,       // use default DB
	})
	str := &redisString{rdb: rdb}
	hash := &redisHash{rdb: rdb}
	return &Client{string: str, hash: hash}
}
