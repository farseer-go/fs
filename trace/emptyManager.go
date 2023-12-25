package trace

type EmptyManager struct {
}

func (m *EmptyManager) EntryWebApi(domain string, path string, method string, contentType string, header map[string]string, requestBody string, requestIp string) ITraceContext {
	return &emptyTraceContext{}
}

func (m *EmptyManager) EntryMqConsumer(server string, queueName string, routingKey string) ITraceContext {
	return &emptyTraceContext{}
}

func (m *EmptyManager) EntryQueueConsumer(subscribeName string) ITraceContext {
	return &emptyTraceContext{}
}

func (m *EmptyManager) EntryTask(taskName string) ITraceContext {
	return &emptyTraceContext{}
}

func (m *EmptyManager) EntryFSchedule(taskGroupName string, taskGroupId int64, taskId int64) ITraceContext {
	return &emptyTraceContext{}
}

func (m *EmptyManager) EntryWatchKey(key string) ITraceContext { return &emptyTraceContext{} }
func (m *EmptyManager) TraceMq(method string, server string, exchange string) ITraceDetail {
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
func (*EmptyManager) TraceHand(name string) ITraceDetail        { return &emptyTraceDetail{} }
func (*EmptyManager) TraceKeyLocation(name string) ITraceDetail { return &emptyTraceDetail{} }
func (*EmptyManager) TraceMqSend(method string, server string, exchange string, routingKey string) ITraceDetail {
	return &emptyTraceDetail{}
}
func (*EmptyManager) TraceRedis(method string, key string, field string) ITraceDetail {
	return &emptyTraceDetail{}
}
func (*EmptyManager) TraceHttp(method string, url string) ITraceDetail { return &emptyTraceDetail{} }

type emptyTraceContext struct{}

func (receiver *emptyTraceContext) Error(err error)                                                 {}
func (receiver *emptyTraceContext) SetBody(requestBody string, statusCode int, responseBody string) {}
func (receiver *emptyTraceContext) GetTraceId() int64                                               { return 0 }
func (receiver *emptyTraceContext) GetStartTs() int64                                               { return 0 }
func (receiver *emptyTraceContext) End()                                                            {}
func (receiver *emptyTraceContext) Ignore()                                                         {}
func (receiver *emptyTraceContext) AddDetail(detail ITraceDetail)                                   {}
func (receiver *emptyTraceContext) GetList() []ITraceDetail {
	return []ITraceDetail{}
}

type emptyTraceDetail struct{}

func (receiver *emptyTraceDetail) GetLevel() int                    { return 0 }
func (receiver *emptyTraceDetail) IsIgnore() bool                   { return true }
func (receiver *emptyTraceDetail) ToString() string                 { return "" }
func (receiver *emptyTraceDetail) GetTraceDetail() *BaseTraceDetail { return &BaseTraceDetail{} }
func (receiver *emptyTraceDetail) End(err error)                    {}
func (receiver *emptyTraceDetail) Ignore()                          {}
func (receiver *emptyTraceDetail) SetSql(connectionString string, DbName string, tableName string, sql string, rowsAffected int64) {
}
