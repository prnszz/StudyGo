// Version: 0.1
// Description: 服务器端
// server.go

package main

import (
	"fmt"
	"net"
)

// Server, 服务器端的结构体, 包含ip地址和端口号, 用于储存服务器端的ip地址和端口号
type Server struct {
	Ip   string
	Port int
}

// 创建一个server的实例
// *Server, 服务器端的结构体指针
func NewServer(ip string, port int) *Server {
	// &Server, 服务器端的结构体指针
	server := &Server{
		Ip:   ip,
		Port: port,
	}

	return server
}

func (this *Server) Handler(conn net.Conn) {
	fmt.Println("connect success")
}

// 启动服务器端口
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port)) // 创建一个监听器, 监听指定的ip地址和端口号
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// close listen socket
	defer listener.Close() // 延迟关闭监听器

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}

		// do handler
		go this.Handler(conn) // 处理连接, 开启一个协程

	}

}
