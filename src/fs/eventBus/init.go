package eventBus

// 订阅者
var subscriber map[string][]ISubscribe

func init() {
	subscriber = make(map[string][]ISubscribe)
}
