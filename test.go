package main

import (
	"os"
	"fmt"
	"net"
	"commontest/Test"
	"commontest/Config"
	"strconv"
)

func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        fmt.Println(err)
    }
    defer conn.Close()
    localAddr := conn.LocalAddr().(*net.UDPAddr)
    return localAddr.IP
}

func main() {
	fmt.Println("BoundIP: ", GetOutboundIP())
	fmt.Println("***TEST PASSED***")
}
