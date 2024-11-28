# Go IM System v0.4
## 业务逻辑变化
将原本在`Server.Handler`中的用户上线, 消息处理, 下线逻辑抽取到`User.Online()`, `User.DoMessage(msg string)`, `User.Offline()`方法中. 

## 代码变化
1. user.go
  - 添加了`server *Server`字段, 使每个User实例都持有对Server的引用.   
  - `NewUser` 函数签名更新, 增加了 `server *Server` 参数

2. server.go
  - Handler 方法简化, 将具体业务逻辑委托给 User 的相应方法处理
  - 创建 User 实例时传入 server 引用: `user := NewUser(conn, this)`
  - 用户上线直接调用 `user.Online()`, 而不是在 Handler 中直接操作 OnlineMap