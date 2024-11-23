# Go IM System

这是一个基于Go语言开发的即时通讯系统的基础框架。

## 项目结构

```
.
├── main.go    // 程序入口文件
└── server.go  // 服务器实现文件
```

## 功能特点

- 基于TCP协议的服务器实现
- 支持并发客户端连接
- 简单的服务器框架设计

## 技术实现

- 使用Go标准库 `net` 包实现TCP服务器
- 采用goroutine处理并发连接
- 使用结构体封装服务器功能

## 快速开始

### 运行服务器

1. 确保已安装Go环境
2. 克隆代码到本地
3. 在项目根目录执行：

```bash
go run main.go server.go
```

服务器将在本地启动，监听8888端口。

### 测试连接

可以使用`nc`测试服务器连接：

```bash
nc 127.0.0.1 8888
```

## 代码结构说明

### main.go

程序的入口文件，创建并启动服务器实例：
```go
func main() {
    server := NewServer("127.0.0.1", 8888)
    server.Start()
}
```

### server.go

服务器的核心实现：

- `Server` 结构体：维护服务器的基本信息
  ```go
  type Server struct {
      Ip   string
      Port int
  }
  ```

- `NewServer` 函数：创建服务器实例
- `Handler` 方法：处理客户端连接
- `Start` 方法：启动服务器并监听连接

