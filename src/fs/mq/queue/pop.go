package queue

import (
	"fs/linq"
	"sync"
	"time"
)

func pop() {
	for true {
		// 遍历队列名称
		for _, queueList := range queueConsumer {
			// 每个队列，使用并行执行
			go queueList.iterationSubscriber()
		}
		time.Sleep(500 * time.Millisecond)
	}
}

// iterationSubscriber 按订阅者遍历
func (queueList *queueList) iterationSubscriber() {
	// 队列没数据时，跳过
	if queueList.queue == nil {
		return
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(queueList.queueSubscribers))

	for _, myQueue := range queueList.queueSubscribers {
		// 得出未消费的长度
		currPullCount := len(queueList.queue) - myQueue.offset - 1

		// 如果未消费的长度小于1，则说明有新的数据需要消费
		if currPullCount < 1 {
			waitGroup.Done()
			continue
		}

		// 计算本次应拉取的数量
		if currPullCount > myQueue.pullCount {
			currPullCount = myQueue.pullCount
		}

		// 得到本次消费的队列切片
		startIndex := myQueue.offset + 1
		endIndex := startIndex + currPullCount
		curQueue := queueList.queue[startIndex:endIndex]
		remainingCount := len(queueList.queue) - endIndex

		// 保存本次消费的位置
		myQueue.offset = endIndex - 1
		// 每个订阅者的消费都是并行的
		go func(myQueue *queueSubscriber, curQueue *[]any, remainingCount int) {
			myQueue.subscriber.Consumer(myQueue.subscribeName, *curQueue, remainingCount)
			waitGroup.Done()
		}(myQueue, &curQueue, remainingCount)
	}

	// 等待所有订阅者并行执行完
	waitGroup.Wait()
	queueList.statLastIndex()
	queueList.curtUsedQueue()
}

// 得到当前所有订阅者的最后消费的位置的最小值
func (queueList *queueList) statLastIndex() {
	queueList.minOffset = linq.Order[*queueSubscriber, int](queueList.queueSubscribers).Min(func(item *queueSubscriber) int {
		return item.offset
	})
}

// 缩减使用过的队列
func (queueList *queueList) curtUsedQueue() {
	// 没有使用，则不缩减
	if queueList.minOffset < 0 {
		return
	}
	// 缩减队列
	queueList.minOffset += 1
	queueList.queue = queueList.queue[queueList.minOffset:]
	// 设置每个订阅者的偏移量
	for _, subscriberQueue := range queueList.queueSubscribers {
		subscriberQueue.offset -= queueList.minOffset
	}
	queueList.minOffset = -1
}
