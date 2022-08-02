package redis

import (
	"encoding/json"
	"github.com/farseernet/farseer.go/linq"
	"github.com/go-redis/redis/v8"
	"reflect"
	"strings"
)

type redisHash struct {
	rdb *redis.Client
}

// SetEntity 添加并序列化成json
func (redisHash *redisHash) SetEntity(key string, field string, entity any) error {
	jsonContent, _ := json.Marshal(entity)
	return redisHash.rdb.HSet(ctx, key, field, jsonContent).Err()
}

// Set 添加
//   - HSet("myhash", "key1", "value1", "key2", "value2")
//   - HSet("myhash", []string{"key1", "value1", "key2", "value2"})
//   - HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
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

// ToEntity 获取单个对象
// 	var client DomainObject
//	_ = repository.Client.Hash.ToEntity("redisKey", "field", &client)
func (redisHash *redisHash) ToEntity(key string, field string, entity any) error {
	jsonContent, _ := redisHash.rdb.HGet(ctx, key, field).Result()
	// 反序列
	return json.Unmarshal([]byte(jsonContent), entity)
}

// GetAll 获取所有集合数据
func (redisHash *redisHash) GetAll(key string) (map[string]string, error) {
	return redisHash.rdb.HGetAll(ctx, key).Result()
}

// ToList 将hash.value反序列化成切片对象
// 	type record struct {
// 		Id int `json:"id"`
// 	}
// 	var records []record
// 	ToList("test", &records)
func (redisHash *redisHash) ToList(key string, arrSlice any) error {
	arrVal := reflect.ValueOf(arrSlice).Elem()
	if arrVal.Kind() != reflect.Slice {
		panic("arr入参必须为切片类型")
	}

	result, err := redisHash.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return err
	}

	// 转成[]string
	arrJson := linq.Map(result).ToValue()
	// 组装成json数组
	jsonContent := "[" + strings.Join(arrJson, ",") + "]"
	// 反序列
	return json.Unmarshal([]byte(jsonContent), arrSlice)
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
