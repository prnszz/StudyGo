// Description: 定义用户结构体及相关方法.
// Version: 0.5
// user.go

package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string // 用于发送数据的管道
	conn net.Conn

	server *Server
}

// 创建一个user的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String() // 获取连接的远程地址

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	// 启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

// 用户上线的业务
func (this *User) Online() {
	//用户上线, 将用户加入到OnlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// 广播当前用户上线消息
	this.server.BroadCast(this, "logged in")
}

// 用户下线的业务
func (this *User) Offline() {
	//用户下线, 将用户从OnlineMap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	// 广播当前用户上线消息
	this.server.BroadCast(this, "logged out")
}

func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg)) // 将消息发送给客户端 '[]'表示byte类型
}

// 用户处理消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "online...\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename:" {
		newName := msg[7:]

		// 判断name是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("The name has been used by others\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("You have updated your name to " + this.Name + "\n")
		}
	} else {
		this.server.BroadCast(this, msg)
	}
}

// 监听当前user channel的方法, 一旦有消息, 就将消息发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C                     // 监听当前user channel的消息
		this.conn.Write([]byte(msg + "\n")) // 将消息发送给客户端
	}
}
