// Description: 定义用户结构体及相关方法.
// Version: 0.2
// user.go

package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string // 用于发送数据的管道
	conn net.Conn
}

// 创建一个user的API
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String() // 获取连接的远程地址

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	// 启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

// 监听当前user channel的方法, 一旦有消息, 就将消息发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C                     // 监听当前user channel的消息
		this.conn.Write([]byte(msg + "\n")) // 将消息发送给客户端
	}
}
