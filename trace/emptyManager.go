package trace

type EmptyManager struct {
}

func (receiver *EmptyManager) Ignore() {}

func (receiver *EmptyManager) GetTraceContext() (*TraceContext, bool) {
	return &TraceContext{}, true
}
func (*EmptyManager) EntryWebApi(domain string, path string, method string, contentType string, header map[string]string, requestIp string) *TraceContext {
	return &TraceContext{List: make([]*TraceDetail, 0)}
}
func (*EmptyManager) EntryWebSocket(domain string, path string, header map[string]string, requestIp string) *TraceContext {
	return &TraceContext{List: make([]*TraceDetail, 0)}
}

func (*EmptyManager) EntryMqConsumer(parentTraceId, parentAppName, server string, queueName string, routingKey string) *TraceContext {
	return &TraceContext{List: make([]*TraceDetail, 0)}
}
func (*EmptyManager) EntryQueueConsumer(queueName, subscribeName string) *TraceContext {
	return &TraceContext{List: make([]*TraceDetail, 0)}
}
func (*EmptyManager) EntryEventConsumer(server, eventName, subscribeName string) *TraceContext {
	return &TraceContext{List: make([]*TraceDetail, 0)}
}
func (*EmptyManager) EntryTask(taskName string) *TraceContext {
	return &TraceContext{List: make([]*TraceDetail, 0)}
}

func (*EmptyManager) EntryFSchedule(taskGroupName string, taskId int64, data map[string]string) *TraceContext {
	return &TraceContext{List: make([]*TraceDetail, 0)}
}
func (*EmptyManager) EntryTaskGroup(taskName string, taskGroupName string, taskId int64) *TraceContext {
	return &TraceContext{List: make([]*TraceDetail, 0)}
}

func (*EmptyManager) EntryWatchKey(key string) *TraceContext {
	return &TraceContext{List: make([]*TraceDetail, 0)}
}
func (*EmptyManager) GetCurTrace() *TraceContext { return nil }
func (*EmptyManager) GetTraceId() string {
	return ""
}
func (*EmptyManager) TraceDatabase() *TraceDetail { return &TraceDetail{} }
func (*EmptyManager) TraceDatabaseOpen(dbName string, connectString string) *TraceDetail {
	return &TraceDetail{}
}
func (*EmptyManager) TraceElasticsearch(method string, IndexName string, AliasesName string) *TraceDetail {
	return &TraceDetail{}
}
func (*EmptyManager) TraceEtcd(method string, key string, leaseID int64) *TraceDetail {
	return &TraceDetail{}
}
func (*EmptyManager) TraceHand(name string) *TraceDetail              { return &TraceDetail{} }
func (*EmptyManager) TraceEventPublish(eventName string) *TraceDetail { return &TraceDetail{} }
func (*EmptyManager) TraceMqSend(method string, server string, exchange string, routingKey string) *TraceDetail {
	return &TraceDetail{}
}
func (*EmptyManager) TraceMq(method string, server string, exchange string) *TraceDetail {
	return &TraceDetail{}
}
func (*EmptyManager) TraceRedis(method string, key string, field string) *TraceDetail {
	return &TraceDetail{}
}
func (*EmptyManager) TraceHttp(method string, url string) *TraceDetail { return &TraceDetail{} }

func (*EmptyManager) Push(traceContext *TraceContext, err error) {}
