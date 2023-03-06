package main

import (
	"log"
	"net"
	"time"
)

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	// 组播地址
	multicastAddr := "232.1.10.100:50000"

	// 本地地址（准备转发的本地 IP 和端口）
	localAddr := "192.168.10.181:2067"

	// 监听组播地址
	udpAddr, err := net.ResolveUDPAddr("udp", multicastAddr)
	if err != nil {
		log.Fatal(err)
	}

	// 指定要转发给的本地地址
	destAddr, err := net.ResolveUDPAddr("udp", localAddr)
	if err != nil {
		log.Fatal(err)
	}

	for i, iface := range ifaces {
		idx := i
		if (iface.Flags&net.FlagUp) == 0 || (iface.Flags&net.FlagMulticast) == 0 {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}

		var okc bool

		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				okc = true
				log.Println("net ", ipnet.IP, " et: ", iface.Name)
			}
		}

		if okc {
			//log.Println("net ", ipnet.IP, " et: ", iface.Name)
			//ipx := ipnet.IP.String()
			ifc := ifaces[idx]
			go func() {
				// 创建连接
				conn, err := net.ListenMulticastUDP("udp", &ifc, udpAddr)
				if err != nil {
					log.Fatal(err)
				}
				log.Println(" et: ", ifc.Name)

				// 循环转发
				//buffer := make([]byte, 65535)
				buffer := make([]byte, 65535)
				for {
					log.Println("in....")
					n, _, err := conn.ReadFromUDP(buffer)
					if err != nil {
						log.Println(err)
						break
					}
					log.Println("write: ", n, " addr: ", ifc.Name, "  : ", "ipx")

					_, err = conn.WriteToUDP(buffer[:n], destAddr)
					if err != nil {
						log.Fatal(err)
					}
				}
			}()
		}
	}

	for {
		time.Sleep(time.Second)
	}
}
