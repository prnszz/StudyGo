# Go IM System v0.5

1. 新增了用户查询功能
   - 添加了一个新的消息处理分支：当用户发送"who"命令时，会返回所有在线用户的信息
   - 在 `DoMessage` 方法中增加了对 "who" 命令的处理逻辑
   ```go
   if msg == "who" {
       this.server.mapLock.Lock()
       for _, user := range this.server.OnlineMap {
           onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "online...\n"
           this.SendMsg(onlineMsg)
       }
       this.server.mapLock.Unlock()
   }
   ```

2. 抽取了消息发送逻辑
   - 新增了 `SendMsg` 方法，将发送消息的功能封装成独立的方法
   - 这样可以统一管理消息发送的格式和行为
   ```go
   func (this *User) SendMsg(msg string) {
       this.conn.Write([]byte(msg))
   }
   ```

