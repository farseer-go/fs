package configure

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("./farseer.yaml")
	err := viper.ReadInConfig() //找到并读取配置文件
	if err != nil {             // 捕获读取中遇到的error
		fmt.Errorf("读取配置文件farseer.yaml时发生错误: %w \n", err)
	}
}

func GetString(key string) string {
	return viper.GetString(key)
}

// SetDefault 设置配置的默认值
func SetDefault(key string, value interface{}) {
	viper.SetDefault(key, value)
}
