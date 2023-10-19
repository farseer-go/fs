package trace

import (
	"github.com/farseer-go/collections"
)

type EmptyManager struct {
}

func (*EmptyManager) GetCurTrace() ITraceContext { return nil }

func (*EmptyManager) TraceWebApi(domain string, path string, method string, contentType string, headerDictionary collections.ReadonlyDictionary[string, string], requestBody string, requestIp string) ITraceContext {
	return &emptyTraceContext{}
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

func (receiver *emptyTraceContext) SetBody(requestBody string, statusCode int, responseBody string) {}
func (receiver *emptyTraceContext) GetTraceId() int64                                               { return 0 }
func (receiver *emptyTraceContext) GetStartTs() int64                                               { return 0 }
func (receiver *emptyTraceContext) End()                                                            {}
func (receiver *emptyTraceContext) AddDetail(detail ITraceDetail)                                   {}
func (receiver *emptyTraceContext) GetList() collections.List[ITraceDetail] {
	return collections.NewList[ITraceDetail]()
}

type emptyTraceDetail struct{}

func (receiver *emptyTraceDetail) ToString() string                                   { return "" }
func (receiver *emptyTraceDetail) GetTraceDetail() *BaseTraceDetail                   { return &BaseTraceDetail{} }
func (receiver *emptyTraceDetail) End(err error)                                      {}
func (receiver *emptyTraceDetail) SetSql(DbName string, tableName string, sql string) {}
