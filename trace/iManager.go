package trace

// IManager 链路追踪管理
type IManager interface {
	// EntryWebApi 创建webapi的链路追踪入口
	EntryWebApi(domain string, path string, method string, contentType string, headerDictionary map[string]string, requestIp string) *TraceContext
	// EntryWebSocket 创建WebSocket的链路追踪入口
	EntryWebSocket(domain string, path string, headerDictionary map[string]string, requestIp string) *TraceContext
	// EntryMqConsumer 创建MQ消费入口
	EntryMqConsumer(parentTraceId, parentAppName, server string, queueName string, routingKey string) *TraceContext
	// EntryQueueConsumer 创建Queue消费入口
	EntryQueueConsumer(queueName, subscribeName string) *TraceContext
	// EntryEventConsumer 创建Event消费入口
	EntryEventConsumer(server, eventName, subscribeName string) *TraceContext
	// EntryTask 创建本地任务入口
	EntryTask(taskName string) *TraceContext
	// EntryTaskGroup 创建本地任务入口（调度中心专用）
	EntryTaskGroup(taskName string, taskGroupName string, taskId int64) *TraceContext
	// EntryFSchedule 创建调度中心入口
	EntryFSchedule(taskGroupName string, taskId int64, data map[string]string) *TraceContext
	// EntryWatchKey 创建etcd入口
	EntryWatchKey(key string) *TraceContext

	// TraceDatabaseOpen 数据库埋点
	TraceDatabaseOpen(dbName string, connectString string) ITraceDetail
	// TraceDatabase 数据库埋点
	TraceDatabase() ITraceDetail
	// TraceElasticsearch Elasticsearch埋点
	TraceElasticsearch(method string, IndexName string, AliasesName string) ITraceDetail
	// TraceEtcd etcd埋点
	TraceEtcd(method string, key string, leaseID int64) ITraceDetail
	// TraceHand 手动埋点
	TraceHand(name string) ITraceDetail
	// TraceEventPublish 事件发布
	TraceEventPublish(eventName string) ITraceDetail
	// TraceMqSend send埋点
	TraceMqSend(method string, server string, exchange string, routingKey string) ITraceDetail
	// TraceMq open、create埋点
	TraceMq(method string, server string, exchange string) ITraceDetail
	// TraceRedis Redis埋点
	TraceRedis(method string, key string, field string) ITraceDetail
	// TraceHttp http埋点
	TraceHttp(method string, url string) ITraceDetail
	// 推送到队列
	Push(traceContext *TraceContext, err error)
}
