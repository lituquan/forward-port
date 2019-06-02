package udp
import (
"fmt"
"net"
	"log"
)

// Start a proxy server listen on fromport
// this proxy will then forward all request from fromport to toport
//
// Notice: a service must has been started on toport
func ProxyStart(fromport, toport int) {
	proxyaddr := fmt.Sprintf(":%d", fromport)
	udp_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1"+proxyaddr)
	//创建监听的地址，并且指定udp协议
	if err != nil {
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	// 监听 客户端连接
	conn, err := net.ListenUDP("udp", udp_addr)
	if err != nil {
		fmt.Println("net.ListenUDP err:", err)
		return
	}
	defer conn.Close()

	for {
		log.Println("start udp")
		buf := make([]byte, 1024)
		n, raddr, err := conn.ReadFromUDP(buf)        //接收客户端发送过来的数据，填充到切片buf中。
		if err != nil {
			return
		}

		log.Println("start udp",string(buf[:n]))
		targetaddr := fmt.Sprintf("127.0.0.1:%d", toport)
		targetconn, err := net.Dial("udp", targetaddr)
		if err != nil {
			fmt.Println("net.Dial err:", err)
			return
		}

		n, err = targetconn.Write(buf[:n])
		if err != nil {
			fmt.Printf("Unable to write to output, error: %s\n", err.Error())
			conn.Close()
			targetconn.Close()
			continue
		}
		go proxyRequest1(targetconn, *conn,raddr)
	}
}
// Forward all requests from r to w
func proxyRequest1(r net.Conn, w net.UDPConn,addr *net.UDPAddr) {
	defer r.Close()
	var buffer = make([]byte, 4096)
	for {
		log.Println("read buffer")
		n, err := r.Read(buffer)
		if err != nil {
			fmt.Printf("Unable to read from input, error: %s\n", err.Error())
			break
		}

		n, err = w.WriteToUDP(buffer[:n],addr)
		if err != nil {
			fmt.Printf("Unable to write to output, error: %s\n", err.Error())
			break
		}
	}
}
