package redis

import (
	"context"
	"github.com/farseernet/farseer.go/configure"
)

type Client struct {
	string *redisString
	hash   *redisHash
}

//redis 客户端
var rdb *redis.Client

//上下文定义
var ctx = context.Background()

//NewClient 初始化
func NewClient(redisName string) *Client {
	configString := configure.GetString("Redis." + redisName)
	if configString == "" {
		panic("[farseer.yaml]找不到相应的配置：Redis." + redisName)
	}
	redisConfig := configure.ParseConfig[redisConfig](configString)
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Server,   //localhost:6379
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DB,       // use default DB
	})
	return &Client{}
}

// Remove 删除
func (redisClient *Client) Remove(keys string) (bool, error) {
	result, err := rdb.Del(ctx, keys).Result()
	return result > 0, err
}

// Exists key值是否存在
func (redisClient *Client) Exists(key string) (bool, error) {
	result, err := rdb.Exists(ctx, key).Result()
	return result > 0, err
}
