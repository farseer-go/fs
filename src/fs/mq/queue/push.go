package queue

import "fs/linq"

// Push 添加数据到队列中
func Push(queueName string, message any) {
	// 首先从订阅者中找到是否存在eventName
	if !linq.Dictionary(queueConsumer).ExistsKey(queueName) {
		panic("未找到队列名称：" + queueName + "，需要先通过订阅队列后，才能Push数据")
	}

	// 添加数据到队列
	queueList := queueConsumer[queueName]
	queueList.queue = append(queueList.queue, message)
}
