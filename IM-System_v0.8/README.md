# Go IM System v0.8

1. 主要改进：在 DoMessage 方法中添加了私聊功能的处理逻辑

```go
else if len(msg) > 4 && msg[:3] == "to|" {
    //1. 获取对方的用户名
    remoteName := strings.Split(msg, "|")[1]
    if remoteName == "" {
        this.SendMsg("The private message format is wrong\nplease use 'to|username|message'\n")
        return
    }

    //2. 根据用户名得到对方User对象
    remoteUser, ok := this.server.OnlineMap[remoteName]
    if !ok {
        this.SendMsg("The user does not exist\n")
        return
    }

    //3. 获取消息内容, 发送消息
    content := strings.Split(msg, "|")[2]
    if content == "" {
        this.SendMsg("The message content is empty\n")
        return
    }
    remoteUser.SendMsg(this.Name + " send a following message to you: " + content + "\n")
}
```

使用的 Go 语法特性：

1. strings 包的使用：
```go
import "strings"
strings.Split(msg, "|")  // 使用Split方法按"|"分割字符串
```

2. 切片操作：
```go
msg[:3] == "to|"  // 获取字符串前3个字符
```

3. 错误处理和输入验证：
```go
if remoteName == "" {
    this.SendMsg("The private message format is wrong\n")
    return
}
```

4. map 查找操作：
```go
remoteUser, ok := this.server.OnlineMap[remoteName]
if !ok {
    this.SendMsg("The user does not exist\n")
    return
}
```
