package trace

type EmptyManager struct {
}

func (*EmptyManager) EntryWebApi(domain string, path string, method string, contentType string, header map[string]string, requestIp string) *TraceContext {
	return &TraceContext{}
}
func (*EmptyManager) EntryWebSocket(domain string, path string, header map[string]string, requestIp string) *TraceContext {
	return &TraceContext{}
}

func (*EmptyManager) EntryMqConsumer(parentTraceId, parentAppName, server string, queueName string, routingKey string) *TraceContext {
	return &TraceContext{}
}
func (*EmptyManager) EntryQueueConsumer(queueName, subscribeName string) *TraceContext {
	return &TraceContext{}
}
func (*EmptyManager) EntryEventConsumer(server, eventName, subscribeName string) *TraceContext {
	return &TraceContext{}
}
func (*EmptyManager) EntryTask(taskName string) *TraceContext {
	return &TraceContext{}
}

func (*EmptyManager) EntryFSchedule(taskGroupName string, taskId int64, data map[string]string) *TraceContext {
	return &TraceContext{}
}
func (*EmptyManager) EntryTaskGroup(taskName string, taskGroupName string, taskId int64) *TraceContext {
	return &TraceContext{}
}

func (*EmptyManager) EntryWatchKey(key string) *TraceContext { return &TraceContext{} }
func (*EmptyManager) GetCurTrace() *TraceContext             { return nil }
func (*EmptyManager) GetTraceId() string {
	return ""
}
func (*EmptyManager) TraceDatabase() ITraceDetail { return &emptyTraceDetail{} }
func (*EmptyManager) TraceDatabaseOpen(dbName string, connectString string) ITraceDetail {
	return &emptyTraceDetail{}
}
func (*EmptyManager) TraceElasticsearch(method string, IndexName string, AliasesName string) ITraceDetail {
	return &emptyTraceDetail{}
}
func (*EmptyManager) TraceEtcd(method string, key string, leaseID int64) ITraceDetail {
	return &emptyTraceDetail{}
}
func (*EmptyManager) TraceHand(name string) ITraceDetail              { return &emptyTraceDetail{} }
func (*EmptyManager) TraceEventPublish(eventName string) ITraceDetail { return &emptyTraceDetail{} }
func (*EmptyManager) TraceMqSend(method string, server string, exchange string, routingKey string) ITraceDetail {
	return &emptyTraceDetail{}
}
func (*EmptyManager) TraceMq(method string, server string, exchange string) ITraceDetail {
	return &emptyTraceDetail{}
}
func (*EmptyManager) TraceRedis(method string, key string, field string) ITraceDetail {
	return &emptyTraceDetail{}
}
func (*EmptyManager) TraceHttp(method string, url string) ITraceDetail { return &emptyTraceDetail{} }

func (*EmptyManager) Push(traceContext *TraceContext, err error) {}

type emptyTraceDetail struct{}

func (*emptyTraceDetail) GetLevel() int                    { return 0 }
func (*emptyTraceDetail) Run(fn func())                    {}
func (*emptyTraceDetail) IsIgnore() bool                   { return true }
func (*emptyTraceDetail) ToString() string                 { return "" }
func (*emptyTraceDetail) GetTraceDetail() *BaseTraceDetail { return &BaseTraceDetail{} }
func (*emptyTraceDetail) End(err error)                    {}
func (*emptyTraceDetail) Ignore()                          {}
func (*emptyTraceDetail) SetSql(connectionString string, DbName string, tableName string, sql string, rowsAffected int64) {
}
func (*emptyTraceDetail) SetHttpRequest(url string, reqHead map[string]any, rspHead map[string]string, requestBody string, responseBody string, statusCode int) {
}

func (*emptyTraceDetail) SetRows(rows int) {
}
