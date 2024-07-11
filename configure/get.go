package configure

import (
	"fmt"
	"github.com/farseer-go/fs/core"
	"github.com/farseer-go/fs/parse"
	"os"
	"strings"
)

// InitConfig 初始化配置文件
func InitConfig() {
	var ymlFile string
	if fsEnv := os.Getenv("fsenv"); fsEnv != "" {
		ymlFile = fmt.Sprintf("./farseer.%s.yaml", fsEnv)
	} else {
		ymlFile = "./farseer.yaml"
	}
	configurationBuilder.AddYamlFile(ymlFile)
	configurationBuilder.AddEnvironmentVariables()
	// 配置文件，我们都是通过a.b访问的。而环境变量是a_b。
	// 让环境变量支持a.b的方式，使用替换的方式以支持。
	configurationBuilder.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// 找到并读取配置文件
	err := configurationBuilder.Build()
	if err != nil { // 捕获读取中遇到的error
		fmt.Printf("An error occurred while reading: %s \n", err)
		return
	}

	// 尝试加载fops中心的配置
	if core.AppName != "fops" && GetFopsServer() != "" {
		lstFopsConfigure, err := getFopsConfigure()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, item := range lstFopsConfigure {
			if item.AppName == "global" {
				item.AppName = ""
			}
		}

		// 重新读取yml文件
		data, _ := os.ReadFile(ymlFile)
		dataContent := string(data)
		for _, fopsConfigure := range lstFopsConfigure {
			dataContent = strings.ReplaceAll(dataContent, "{{"+fopsConfigure.Key+"}}", fopsConfigure.Value)
		}

		// 重新写入yml提供者
		if err = ymlProvider.LoadContent([]byte(dataContent)); err != nil {
			fmt.Printf("There is a problem with the configuration read through the configuration center: %s \n", err)
			return
		}
	}
}

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
