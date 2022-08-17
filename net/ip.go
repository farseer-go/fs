package net

// 获取本机IP地址
import "net"

var Ip string
var Ips []string

func InitNet() {
	Ips, _ = LocalIPv4s()
	Ip = Ips[0]
}

// LocalIPv4s 获取本机IP地址
func LocalIPv4s() ([]string, error) {
	var ips []string
	address, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, a := range address {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	return ips, nil
}
