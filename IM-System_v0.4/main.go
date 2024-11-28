// Description: Main entry of the IM system.
// main.go
// version: 0.4

package main

func main() {
	server := NewServer("127.0.0.1", 8888) // 创建一个server的实例, ip地址是本地主机, 端口号是8888
	server.Start()                         // 启动服务器端口
}
