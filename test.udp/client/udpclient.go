package main

import (
	"net"
	"fmt"
	"log"
	"time"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	conn.Write([]byte(time.Now().String()))
	var ch=make(chan []byte)
	go func() {
		read := make([]byte, 4096)
		n, err := conn.Read(read)
		if err != nil {
			log.Println("read err", err)
			return
		}
		ch<-read[:n]
	}()
	log.Println( conn.RemoteAddr().String(),string(<-ch))
}