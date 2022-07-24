### eventbus使用场景
以事件驱动的方式来解耦业务逻辑，在`DDD`中，事件总线是必然用到的技术。

当两个业务模块相互之间有业务关联，但又不希望在代码结构上直接依赖。则可以使用事件驱动的方式来解耦相互之间的依赖。

比如，当一个用户注册成功后，你希望：
1. 统计每天的用户注册量；
2. 发送一条手机短信到用户手机；
3. 将感兴趣的话题推送给当前用户

可以看出这3个动作，并非用户的职责。我们也不希望在用户的注册方法函数中，去依赖手机短信、邮件和话题推送的任何代码。

### 如何使用

我们可以利用事件总线来实现代码的解耦：

有三个订阅者（他们在各自的包中），分别是
1. statUserConsumer
2. newUserConsumer
3. topicPushConsumer

**导入包：import "fs/eventBus"**

```go
// 统计每天的用户注册量
func statUserConsumer(message any, ea EventArgs) {
    user := message.(NewUser)
    // do.....
}

// 发送一条手机短信到用户手机
func newUserConsumer(message any, ea EventArgs) {
    user := message.(NewUser)
    // do.....
}

// 将用户感兴趣的话题推送给当前用户
func topicPushConsumer(message any, ea EventArgs) {
    user := message.(NewUser)
    // do.....
}
```

```go
// 订阅new_user_event事件
eventBus.Subscribe("new_user_event", statUserConsumer)
eventBus.Subscribe("new_user_event", newUserConsumer)
eventBus.Subscribe("new_user_event", topicPushConsumer)
```

然后在新用户注册时，发布事件
```go
type NewUser struct {
    UserName string
}

func UserRegister(userName string,pwd string) {
    eventBus.PublishEvent("new_user_event", newUser{UserName: userName})
}
```

### 同步、异步
有两种事件发布方式，同步（阻塞）或异步的方式。

同步发送时，只有当订阅方完成调用后，才能执行后续的代码

异步发送时，不会阻塞，调用`PublishEventAsync`后，立即返回
```go
// 同步
eventBus.PublishEvent("new_user_event", newUser{UserName: userName})
// 异步
eventBus.PublishEventAsync("new_user_event", newUser{UserName: userName})
```