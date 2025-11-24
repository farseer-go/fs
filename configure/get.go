package configure

import (
	"strings"

	"github.com/farseer-go/fs/parse"
)

// GetString 获取配置
func GetString(key string) string {
	return parse.Convert(configurationBuilder.Get(key), "")
}

// GetStrings 获取配置
func GetStrings(key string) []string {
	val := GetString(key)
	if val == "" {
		return nil
	}
	return strings.Split(val, ",")
}

// GetInt 获取配置
func GetInt(key string) int {
	return parse.Convert(configurationBuilder.Get(key), 0)
}

// GetInt64 获取配置
func GetInt64(key string) int64 {
	return parse.Convert(configurationBuilder.Get(key), int64(0))
}

// GetBool 获取配置
func GetBool(key string) bool {
	return parse.Convert(configurationBuilder.Get(key), false)
}

// GetSubNodes 获取所有子节点
func GetSubNodes(key string) map[string]any {
	return configurationBuilder.GetSubNodes(key)
}

// GetSliceNodes 获取数组节点
func GetSliceNodes(key string) []map[string]any {
	return configurationBuilder.GetSliceNodes(key)
}

// GetSlice 获取数组
func GetSlice(key string) []string {
	return configurationBuilder.GetSlice(key)
}

// SetDefault 设置配置的默认值
func SetDefault(key string, value any) {
	configurationBuilder.def[key] = value
}
