// Version: 0.5
// Description: 服务器端
// server.go

package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

// Server, 服务器端的结构体, 包含ip地址和端口号, 用于储存服务器端的ip地址和端口号
type Server struct {
	Ip   string
	Port int

	// Users, 用于储存所有的用户
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

// 创建一个server的实例
// *Server, 服务器端的结构体指针
func NewServer(ip string, port int) *Server {
	// &Server, 服务器端的结构体指针
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

// 监听Message广播消息channel的goroutine, 一旦有消息就发送给全部的在线user
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message // 监听Message广播消息channel的消息

		// 将msg发送给全部的在线user
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	// fmt.Println("connect success")

	user := NewUser(conn, this) // 创建一个user的实例

	user.Online() // 用户上线

	// 接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf) // 读取客户端发送的消息
			if n == 0 {              // 客户端断开连接
				user.Offline() // 广播用户下线消息
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				user.Offline() // 发生错误时也需要让用户下线
				return
			}

			// 提取用户消息, 去除'\n'
			msg := string(buf[:n-1])

			// 用户针对msg进行消息处理
			user.DoMessage(msg)
		}
	}()

	//当前handler阻塞, 监听用户发送的消息
	select {}
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

	// 启动监听Message广播消息的goroutine
	go this.ListenMessage()

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
