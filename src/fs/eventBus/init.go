package eventBus

// 订阅者
var subscriber map[string][]consumerFunc

func init() {
	subscriber = make(map[string][]consumerFunc)
}
