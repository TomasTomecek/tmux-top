package main

import (
	"fmt"
	"net"
)

func GetNetStats() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	fmt.Println(addrs)
	ifs, err := net.Interfaces()
	if err != nil {
		return
	}
	fmt.Println(ifs)
}

func main() {
	GetNetStats()
}
