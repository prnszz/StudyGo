# Go IM System v0.2

这是一个基于Go语言开发的简单即时通讯系统,用于学习和复习Go语言的核心特性。

## 项目结构

```
.
├── main.go    // 程序入口文件
├── server.go  // 服务器实现文件
└── user.go    // 用户相关实现
```

## 新增功能 (v0.2)

- 用户系统
  - 用户结构体定义
  - 用户上线/下线管理
  - 用户消息接收和发送
- 消息广播系统
  - 服务器消息广播channel
  - 用户消息监听

## 技术要点

本项目用于复习以下Go语言特性：

1. 结构体和方法
```go
type User struct {
    Name string
    Addr string
    C    chan string
    conn net.Conn
}
```

2. goroutine并发
```go
go this.ListenMessage()
go this.Handler(conn)
```

3. channel通信
```go
msg := <-this.Message  // 监听消息
cli.C <- msg          // 发送消息
```

4. 互斥锁(sync.RWMutex)
```go
this.mapLock.Lock()
this.OnlineMap[user.Name] = user
this.mapLock.Unlock()
```

5. 接口使用
```go
conn net.Conn
```

## 快速开始

1. 运行服务器:
```bash
go run *.go
```

2. 使用telnet测试:
```bash
nc 127.0.0.1 8888
```

## 代码结构说明

### server.go
- `Server`结构体：服务器核心结构
  - `OnlineMap`: 在线用户管理
  - `Message`: 广播消息channel
  - `mapLock`: 并发安全锁
- 主要方法：
  - `ListenMessage()`: 监听广播消息
  - `BroadCast()`: 广播消息
  - `Handler()`: 处理用户连接

### user.go
- `User`结构体：用户信息管理
  - 基本信息(Name, Addr)
  - 消息channel
  - 网络连接
- 主要方法：
  - `NewUser()`: 创建用户
  - `ListenMessage()`: 监听用户消息
