package eventBus

type consumerFunc func(message any, ea EventArgs)

/*type IEventSubscribe interface {
	// Consumer 消费
	Consumer(message any, ea EventArgs)
}*/

// Subscribe 订阅事件
func Subscribe(eventName string, fn consumerFunc) {
	subscriber[eventName] = append(subscriber[eventName], fn)
}
