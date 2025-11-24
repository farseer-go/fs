package configure

import (
	"fmt"
	"os"
	"strings"

	"github.com/farseer-go/fs/core"
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
