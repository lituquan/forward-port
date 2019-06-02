package main

import (
	"net"
	"log"
	"time"
)

func main() {
	// 创建 服务器 UDP 地址结构。指定 IP + port
	log.Println("start")
	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8081")
	if err != nil {
		log.Println("ResolveUDPAddr err:", err)
		return
	}
	// 监听 客户端连接
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Println("net.ListenUDP err:", err)
		return
	}
	defer conn.Close()

	for {
		log.Println("client conn")
		buf := make([]byte, 1024)
		n, raddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("conn.ReadFromUDP err:", err)
			return
		}
		messsage:=buf[:n]
		log.Println("客户端发来信息",string(messsage))
		conn.WriteToUDP([]byte("OK!"+time.Now().String()),raddr)
	}
}