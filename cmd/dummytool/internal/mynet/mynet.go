package mynet

import (
	"net"
	"strings"
)

func ValidateNodeAddr(nodeAddr string) (domain string, ip net.IP, valid bool) {

	if strings.Contains(nodeAddr, "/") {
		parts := strings.Split(nodeAddr, "/")
		if len(parts) != 2 {
			return "", nil, false
		}
		domain = parts[0]
		ip = net.ParseIP(parts[1])
		if ip == nil {
			return "", nil, false
		}
		return
	}

	ip = net.ParseIP(nodeAddr)
	valid = ip != nil
	return
}

func CheckIPString(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil
}
