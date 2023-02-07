package core

// EventArgs 事件属性
type EventArgs struct {
	// 唯一标识
	Id string
	// 事件的发布时间
	CreateAt int64
	// 消息内容
	Message any
	// 执行失败次数
	ErrorCount int
	// 事件名称
	EventName string
}

// ConsumerFunc 订阅者的函数
type ConsumerFunc func(message any, ea EventArgs)
