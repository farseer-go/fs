package queue

import "time"

func pop() {
	for true {
		// 遍历队列名称
		for _, queueList := range queueConsumer {
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
					subscriberQueue.lastConsumerIndex = endIndex
					// 消费
					go subscriberQueue.subscriber.Consumer(curQueue)
				}
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}
