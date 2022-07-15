package queue

import "fs/linq"

type IQueueSubscribe interface {
	// Consumer 消费
	Consumer(subscribeName string, message []any, remainingCount int)
}

// Subscribe 订阅
// queueName = 队列名称
// subscribeName = 订阅者名称
// pullCount = 每次拉取的数量
func Subscribe(queueName string, subscribeName string, pullCount int, fn IQueueSubscribe) {
	if !linq.Dictionary(queueConsumer).ExistsKey(queueName) {
		queueConsumer[queueName] = &queueList{
			queueName:        queueName,
			queue:            nil,
			minOffset:        -1,
			queueSubscribers: nil,
		}
	}

	// 找到对应的队列
	queueList := queueConsumer[queueName]
	// 添加订阅者
	queueList.queueSubscribers = append(queueList.queueSubscribers, &queueSubscriber{
		subscribeName: subscribeName,
		offset:        -1,
		subscriber:    fn,
		pullCount:     pullCount,
	})
}
