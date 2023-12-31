package trace

// IManager 链路追踪管理
type IManager interface {
	GetCurTrace() ITraceContext
	// EntryWebApi 创建webapi的链路追踪入口
	EntryWebApi(domain string, path string, method string, contentType string, headerDictionary map[string]string, requestBody string, requestIp string) ITraceContext
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
	// TraceMqSend send埋点
	TraceMqSend(method string, server string, exchange string, routingKey string) ITraceDetail
	// TraceMq open、create埋点
	TraceMq(method string, server string, exchange string) ITraceDetail
	// EntryMqConsumer 创建MQ消费入口
	EntryMqConsumer(server string, queueName string, routingKey string) ITraceContext
	// EntryQueueConsumer 创建Queue消费入口
	EntryQueueConsumer(subscribeName string) ITraceContext
	// TraceRedis Redis埋点
	TraceRedis(method string, key string, field string) ITraceDetail
	// TraceHttp http埋点
	TraceHttp(method string, url string) ITraceDetail
	// EntryTask 创建本地任务入口
	EntryTask(taskName string) ITraceContext
	// EntryTaskGroup 创建本地任务入口（调度中心专用）
	EntryTaskGroup(taskName string, taskGroupName string, taskGroupId int64, taskId int64) ITraceContext
	// EntryFSchedule 创建调度中心入口
	EntryFSchedule(taskGroupName string, taskGroupId int64, taskId int64) ITraceContext
	// EntryWatchKey 创建etcd入口
	EntryWatchKey(key string) ITraceContext
}
