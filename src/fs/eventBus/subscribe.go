package eventBus

type ISubscribe interface {
	// Consumer 消费
	Consumer(message any, ea EventArgs)
}

// Subscribe 订阅
func Subscribe(eventName string, fn ISubscribe) {
	subscriber[eventName] = append(subscriber[eventName], fn)
}
