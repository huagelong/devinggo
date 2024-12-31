// Package validate
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package validate

import "net"

func IsIp(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IsPublicIp 是否是公网IP
func IsPublicIp(ip string) bool {
	i := net.ParseIP(ip)

	if i.IsLoopback() || i.IsPrivate() || i.IsMulticast() || i.IsUnspecified() || i.IsLinkLocalUnicast() || i.IsLinkLocalMulticast() {
		return false
	}

	if ip4 := i.To4(); ip4 != nil {
		return !i.Equal(net.IPv4bcast)
	}
	return true
}

// IsLocalIPAddr 检测 IP 地址字符串是否是内网地址
func IsLocalIPAddr(ip string) bool {
	if "localhost" == ip {
		return true
	}
	return HasLocalIP(net.ParseIP(ip))
}

// HasLocalIP 检测 IP 地址是否是内网地址
func HasLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}
