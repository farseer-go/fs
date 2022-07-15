package queue

func init() {
	queueConsumer = make(map[string]*queueList)
	// 启动消费
	go pop()
}
