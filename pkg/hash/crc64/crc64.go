package main

import (
	"fmt"
	// "hash/crc64"
	"net"
	// "math"
)

func main() {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Error : " + err.Error())
	}
	bs := ""
	for _, inter := range interfaces {
		mac := inter.HardwareAddr //获取本机MAC地址
		bs += fmt.Sprintf("%02X", mac)
	}

	fmt.Println(bs)
	// sum := crc64.Checksum([]byte(bs), crc64.ISO)
	// fmt.Println(sum)

}
