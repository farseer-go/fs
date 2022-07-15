package eventBus

type IEventSubscribe interface {
	// Consumer 消费
	Consumer(message any, ea EventArgs)
}

// Subscribe 订阅
func Subscribe(eventName string, fn IEventSubscribe) {
	subscriber[eventName] = append(subscriber[eventName], fn)
}
