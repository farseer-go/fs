package configure

import (
	"github.com/farseer-go/fs/parse"
	"strings"
)

func ReadInConfig() error {
	configurationBuilder.AddYamlFile("./farseer.yaml")
	configurationBuilder.AddEnvironmentVariables()
	// 配置文件，我们都是通过a.b访问的。而环境变量是a_b。
	// 让环境变量支持a.b的方式，使用替换的方式以支持。
	configurationBuilder.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// 找到并读取配置文件
	return configurationBuilder.Build()
}

// GetString 获取配置
func GetString(key string) string {
	return parse.Convert(configurationBuilder.Get(key), "")
}

// GetStrings 获取配置
func GetStrings(key string) []string {
	return strings.Split(GetString(key), ",")
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
