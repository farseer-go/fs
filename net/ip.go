package net

// 获取本机IP地址
import "net"

var ip string
var ips []string

func GetIp() string {
	if ip == "" {
		ips, _ = LocalIPv4s()
		ip = ips[0]
	}
	return ip
}

// LocalIPv4s 获取本机IP地址
func LocalIPv4s() ([]string, error) {
	var ips []string
	address, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, a := range address {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP.String())
		}
	}

	return ips, nil
}
