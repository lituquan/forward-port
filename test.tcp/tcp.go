

package main

import (
"net"
"log"
	"time"
)

func Handle(conn net.Conn)  {
	buf:=make([]byte,4096)
	n,err:=conn.Read(buf)
	if err!=nil{
		log.Println("read err",err)
		return
	}
	messsage:=buf[:n]
	log.Println("客户端发来信息",string(messsage))
	conn.Write([]byte("OK!"+time.Now().String()))
}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Println("listen err", err)
		return
	}
	defer listen.Close()
	for {
		conn,err:=listen.Accept()
		if err!=nil{
			log.Println("accept err",err)
			return
		}
		go Handle(conn)
	}
}

