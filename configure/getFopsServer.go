package configure

import "strings"

func GetFopsServer() string {
	fopsServer := strings.ToLower(GetString("Fops.Server"))
	if !strings.HasPrefix(fopsServer, "http") {
		panic("[farseer.yaml]Fops.Server配置不正确，示例：https://fops.fsgit.com")
	}
	if !strings.HasSuffix(fopsServer, "/") {
		fopsServer += "/"
	}
	return fopsServer
}
