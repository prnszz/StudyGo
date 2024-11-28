# Go IM System v0.3
这个版本主要更新了消息广播 (BroadCast) 方法的实现. 
### 消息广播功能说明
- BroadCast方法实现
  - 格式化消息：`[addr]username:message`
  - 通过Server的Message channel发送
  - 支持的广播类型：
    - 用户上线通知
    - 用户下线通知
    - 用户聊天消息

- Handler中的广播处理
  - 用户上线时广播通知
  - 用户断开连接时广播下线通知
  - 接收用户消息并广播
  - 使用goroutine处理用户消息接收