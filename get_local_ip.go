package main

import (
	"net"
	"strings"
)

type NoValidLocalIPError struct{}

func (p *NoValidLocalIPError) Error() string {
	return "No valid local IP address found."
}

func getLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		if strings.HasPrefix(address.String(), "192.168.") {
			return strings.Split(address.String(), "/")[0], nil
		}
	}
	return "", &NoValidLocalIPError{}
}
