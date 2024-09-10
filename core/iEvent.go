package core

// IEvent 定义的一个通用的事件发布接口
type IEvent interface {
	// Publish 发布消息
	Publish(message any) error
	// PublishAsync 异步发布消息
	PublishAsync(message any)
}
