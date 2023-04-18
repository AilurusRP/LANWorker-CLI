package info

import (
	"net"
	"strings"
)

var IP string

func GetLocalIP() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, address := range addrs {
		if strings.HasPrefix(address.String(), "192.168.") {
			IP = strings.Split(address.String(), "/")[0]
		}
	}
	if IP == "" {
		panic("No valid local IP address found. Please check your network connection.")
	}
}
