package main

import (
	"fmt"
	"net"
	"time"
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "Unknown"
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "Unknown"
}

func main() {
	fmt.Println("Xin chào, thế giới!")
	fmt.Println("Ngày và giờ hiện tại:", time.Now().Format(time.RFC1123))
	fmt.Println("Địa chỉ IP của máy tính:", getLocalIP())
	fmt
}