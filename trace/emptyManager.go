package trace

type EmptyManager struct {
}

func (*EmptyManager) EntryWebApi(domain string, path string, method string, contentType string, header map[string]string, requestBody string, requestIp string) ITraceContext {
	return &emptyTraceContext{}
}

func (*EmptyManager) EntryMqConsumer(server string, queueName string, routingKey string) ITraceContext {
	return &emptyTraceContext{}
}

func (*EmptyManager) EntryQueueConsumer(queueName, subscribeName string) ITraceContext {
	return &emptyTraceContext{}
}

func (*EmptyManager) EntryTask(taskName string) ITraceContext {
	return &emptyTraceContext{}
}

func (*EmptyManager) EntryFSchedule(taskGroupName string, taskId int64, data map[string]string) ITraceContext {
	return &emptyTraceContext{}
}
func (*EmptyManager) EntryTaskGroup(taskName string, taskGroupName string, taskId int64) ITraceContext {
	return &emptyTraceContext{}
}

func (*EmptyManager) EntryWatchKey(key string) ITraceContext { return &emptyTraceContext{} }
func (*EmptyManager) TraceMq(method string, server string, exchange string) ITraceDetail {
	return &emptyTraceDetail{}
}
func (*EmptyManager) GetCurTrace() ITraceContext  { return nil }
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
func (*EmptyManager) TraceHand(name string) ITraceDetail { return &emptyTraceDetail{} }
func (*EmptyManager) TraceMqSend(method string, server string, exchange string, routingKey string) ITraceDetail {
	return &emptyTraceDetail{}
}
func (*EmptyManager) TraceRedis(method string, key string, field string) ITraceDetail {
	return &emptyTraceDetail{}
}
func (*EmptyManager) TraceHttp(method string, url string) ITraceDetail { return &emptyTraceDetail{} }

type emptyTraceContext struct{}

func (*emptyTraceContext) Error(err error)                                                 {}
func (*emptyTraceContext) SetBody(requestBody string, statusCode int, responseBody string) {}
func (*emptyTraceContext) GetTraceId() string                                              { return "" }
func (*emptyTraceContext) GetTraceLevel() int                                              { return 0 }
func (*emptyTraceContext) GetStartTs() int64                                               { return 0 }
func (*emptyTraceContext) End()                                                            {}
func (*emptyTraceContext) Ignore()                                                         {}
func (*emptyTraceContext) AddDetail(detail ITraceDetail)                                   {}
func (*emptyTraceContext) GetList() []any {
	return []any{}
}
func (*emptyTraceContext) GetAppInfo() (string, string, string, string, string) {
	return "", "", "", "", ""
}

type emptyTraceDetail struct{}

func (*emptyTraceDetail) GetLevel() int                    { return 0 }
func (*emptyTraceDetail) IsIgnore() bool                   { return true }
func (*emptyTraceDetail) ToString() string                 { return "" }
func (*emptyTraceDetail) GetTraceDetail() *BaseTraceDetail { return &BaseTraceDetail{} }
func (*emptyTraceDetail) End(err error)                    {}
func (*emptyTraceDetail) Ignore()                          {}
func (*emptyTraceDetail) SetSql(connectionString string, DbName string, tableName string, sql string, rowsAffected int64) {
}
func (*emptyTraceDetail) SetHttpRequest(url string, head map[string]any, requestBody string, responseBody string, statusCode int) {
}
