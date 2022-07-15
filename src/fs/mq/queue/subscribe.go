package queue

import "fs/linq"

type IQueueSubscribe interface {
	// Consumer 消费
	Consumer(message []any)
}

// Subscribe 订阅
// queueName = 队列名称
// pullCount = 每次拉取的数量
func Subscribe(queueName string, pullCount int, fn IQueueSubscribe) {
	if !linq.Dictionary(queueConsumer).ExistsKey(queueName) {
		queueConsumer[queueName] = queueList{
			queueName:        queueName,
			queue:            nil,
			consumerIndex:    -1,
			subscriberQueues: nil,
		}
	}

	// 找到对应的队列
	queueList := queueConsumer[queueName]
	// 添加订阅者
	queueList.subscriberQueues = append(queueList.subscriberQueues, subscriberQueue{
		lastConsumerIndex: -1,
		subscriber:        fn,
		pullCount:         pullCount,
	})
}
