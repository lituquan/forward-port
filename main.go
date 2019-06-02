package main

import (
	"forward-port/tcp"
	"forward-port/udp"
)
var ch=make(chan int)

func main() {
	go tcp.ProxyStart(8080,"127.0.0.1:8081")
	go udp.ProxyStart(8080,"127.0.0.1:8081")
	<-ch
}