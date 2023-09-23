package net

// 获取本机IP地址
import (
	"net"
	"strings"
)

var ip string
var ipList []string

func GetIp() string {
	if ip == "" {
		ipList, _ = LocalIPv4s()
		// 优先使用192.168开头
		for _, strIp := range ipList {
			if strings.HasPrefix(strIp, "192.168.") {
				ip = strIp
				return ip
			}
		}

		// 使用172.20.开头
		for _, strIp := range ipList {
			if strings.HasPrefix(strIp, "172.20.") {
				ip = strIp
				return ip
			}
		}

		// 使用10.开头
		for _, strIp := range ipList {
			if strings.HasPrefix(strIp, "10.") {
				ip = strIp
				return ip
			}
		}

		// 没有192.168.x.x和10.x.x.x开头的，只能取第一个IP返回
		ip = ipList[0]
	}
	return ip
}

// LocalIPv4s 获取本机IP地址
func LocalIPv4s() ([]string, error) {
	var ips []string
	address, err := net.InterfaceAddrs()

	for _, a := range address {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP.String())
		}
	}

	return ips, err
}
