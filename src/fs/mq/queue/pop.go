package queue

import (
	"fs/linq"
	"time"
)

func pop() {
	for true {
		// 遍历队列名称
		for _, queueList := range queueConsumer {
			// 队列没数据时，跳过
			if queueList.queue == nil {
				continue
			}

			// 按订阅者遍历
			for _, subscriberQueue := range queueList.subscriberQueues {
				// 得出未消费的长度
				currPullCount := len(queueList.queue) - subscriberQueue.lastConsumerIndex - 1
				// 如果未消费的长度大于0，则说明有新的数据需要消费
				if currPullCount > 0 {
					// 计算本次应拉取的数量
					if currPullCount > subscriberQueue.pullCount {
						currPullCount = subscriberQueue.pullCount
					}

					// 得到本次消费的队列切片
					startIndex := subscriberQueue.lastConsumerIndex + 1
					endIndex := startIndex + currPullCount
					curQueue := queueList.queue[startIndex:endIndex]

					// 保存本次消费的位置
					subscriberQueue.lastConsumerIndex = endIndex - 1
					// 消费
					go subscriberQueue.subscriber.Consumer(curQueue)
				}
			}

			// 得到当前所有订阅者的最后消费的位置的最小值
			queueList.consumerLastIndex = linq.Order[*subscriberQueue, int](queueList.subscriberQueues).Min(func(item *subscriberQueue) int {
				return item.lastConsumerIndex
			})
		}
		time.Sleep(500 * time.Millisecond)
	}
}
