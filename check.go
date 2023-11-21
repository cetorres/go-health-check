package main

import (
	"fmt"
	"net"
	"time"
)

func Check(domain string, port string) string {
	address := domain + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string
	colorRed := "\033[31m"
  colorGreen := "\033[32m"

	if err != nil {
		status = fmt.Sprintf("%v[DOWN] %v is unreachable.", string(colorRed), domain)
	} else {
		status = fmt.Sprintf("%v[UP] %v (%v) is reachable.", string(colorGreen), domain, conn.RemoteAddr().String())
	}

	return status
}