package mynet

import (
	"net"
)

func CheckIPString(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil
}