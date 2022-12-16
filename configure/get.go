package configure

import (
	"github.com/farseer-go/fs/parse"
	"strings"
)

func ReadInConfig() error {
	configurationBuilder = NewConfigurationBuilder()
	configurationBuilder.AddYamlFile("./farseer.yaml")
	configurationBuilder.AddEnvironmentVariables()
	// 配置文件，我们都是通过a.b访问的。而环境变量是A_B。
	// 让环境变量支持A.B的方式，使用替换的方式以支持。
	configurationBuilder.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// 找到并读取配置文件
	return configurationBuilder.Build()
}

// GetString 获取配置
func GetString(key string) string {
	return configurationBuilder.GetString(key)
}

// GetStrings 获取配置
func GetStrings(key string) []string {
	return strings.Split(GetString(key), ",")
}

// GetInt 获取配置
func GetInt(key string) int {
	return parse.Convert(configurationBuilder.GetString(key), 0)
}

// GetInt64 获取配置
func GetInt64(key string) int64 {
	return parse.Convert(configurationBuilder.GetString(key), int64(0))
}

// GetBool 获取配置
func GetBool(key string) bool {
	return parse.Convert(configurationBuilder.GetString(key), false)
}

// GetSubNodes 获取所有子节点
func GetSubNodes(key string) map[string]any {
	return configurationBuilder.GetSubNodes(key)
}

// GetSlice 获取数组
func GetSlice(key string) []string {
	return configurationBuilder.GetSlice(key)
}

// SetDefault 设置配置的默认值
func SetDefault(key string, value any) {
	configurationBuilder.def[key] = value
}
