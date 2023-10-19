package trace

import (
	"github.com/farseer-go/collections"
)

// IManager 链路追踪管理
type IManager interface {
	GetCurTrace() ITraceContext
	// EntryWebApi 创建webapi的链路追踪入口
	EntryWebApi(domain string, path string, method string, contentType string, headerDictionary collections.ReadonlyDictionary[string, string], requestBody string, requestIp string) ITraceContext
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
	// TraceKeyLocation 关键位置埋点
	TraceKeyLocation(name string) ITraceDetail
	// TraceMqSend send埋点
	TraceMqSend(method string, server string, exchange string, routingKey string) ITraceDetail
	// TraceMq open、create埋点
	TraceMq(method string, server string, exchange string) ITraceDetail
	// EntryMqConsumer 创建MQ消费入口
	EntryMqConsumer(server string, queueName string, routingKey string) ITraceContext
	// TraceRedis Redis埋点
	TraceRedis(method string, key string, field string) ITraceDetail
	// TraceHttp http埋点
	TraceHttp(method string, url string) ITraceDetail
}
