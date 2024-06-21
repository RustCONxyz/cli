package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func ParseConnectionDetails(connectionDetails string) (string, int, error) {
	parts := strings.Split(connectionDetails, ":")
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("invalid connection details format")
	}

	host := parts[0]
	if host != "localhost" && net.ParseIP(host) == nil {
		return "", 0, fmt.Errorf("invalid IP address")
	}

	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, fmt.Errorf("invalid port number")
	}
	if port < 1 || port > 65535 {
		return "", 0, fmt.Errorf("port out of range")
	}

	return host, port, nil
}
