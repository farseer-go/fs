package configure

import (
	"github.com/spf13/viper"
	"strings"
)

func ReadInConfig() error {
	viper.SetConfigFile("./farseer.yaml")
	return viper.ReadInConfig() //找到并读取配置文件
}

// GetString 获取配置
func GetString(key string) string {
	return viper.GetString(key)
}

// GetStrings 获取配置
func GetStrings(key string) []string {
	return strings.Split(GetString(key), ",")
}

// GetSlice 获取数组
func GetSlice(key string) []string {
	return viper.GetStringSlice(key)
}

// GetMap 读取子节点
func GetMap(key string) map[string]string {
	return viper.GetStringMapString(key)
}

// GetInt 获取配置
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetInt64 获取配置
func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

// GetBool 获取配置
func GetBool(key string) bool {
	return viper.GetBool(key)
}

// GetSubNodes 获取所有子节点
func GetSubNodes(key string) map[string]string {
	return viper.GetStringMapString(key)
}

// SetDefault 设置配置的默认值
func SetDefault(key string, value interface{}) {
	viper.SetDefault(key, value)
}
