package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "232.1.10.100:50000")
	if err != nil {
		log.Fatalln(err)
	}

	// 创建UDP监听器
	conn, err := net.ListenPacket("udp", "232.1.10.100:50000")
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Listening on 232.1.10.100:50000...")

	// 读取UDP数据包并处理
	buffer := make([]byte, 65535)
	conn.WriteTo([]byte("1"), udpAddr)
	for {
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}
		fmt.Printf("Received %d bytes from %s:\n%s\n", n, addr.String(), string(buffer[:n]))
	}
}
