package queue

// 队列列表
type queueList struct {
	// 当前队列名称
	queueName string
	// 全局队列
	queue []any
	// 当前消费到的索引位置（如果是多个消费者，只记录最早的索引位置）
	// 用于定时移除queue已被消费的数据，以节省内存空间
	consumerLastIndex int
	// 订阅者
	subscriberQueues []*subscriberQueue
}

// 订阅者的队列
type subscriberQueue struct {
	// 最后消费的位置
	lastConsumerIndex int
	// 订阅者
	subscriber IQueueSubscribe
	// 每次拉取的数量
	pullCount int
}

// 队列
// key = queueName
// value = 队列
var queueConsumer map[string]*queueList
