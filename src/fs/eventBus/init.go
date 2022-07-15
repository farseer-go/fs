package eventBus

// 订阅者
var subscriber map[string][]IEventSubscribe

func init() {
	subscriber = make(map[string][]IEventSubscribe)
}
