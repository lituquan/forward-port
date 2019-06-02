package main

import (
	"net"
	"log"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println("dial err", err)
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