package configure

import (
	"strings"
)

// Fops服务地址
var fopsServer string

// GetFopsServer 获取FOPS地址
func GetFopsServer() string {
	if fopsServer != "" {
		return fopsServer
	}

	if fopsServer = strings.ToLower(GetString("Fops.Server")); fopsServer != "" {
		if !strings.HasPrefix(fopsServer, "http") {
			panic("[farseer.yaml]Fops.Server配置不正确，示例：https://fops.fsgit.com")
		}
		if !strings.HasSuffix(fopsServer, "/") {
			fopsServer += "/"
		}
	}
	return fopsServer
}
