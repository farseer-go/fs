package trace

type ITraceDetail interface {
	// Run 运行完后调用End
	Run(fn func())
	ToString() string
	GetTraceDetail() *BaseTraceDetail
	End(err error)
	SetSql(connectionString string, DbName string, tableName string, sql string, rowsAffected int64)
	// Ignore 忽略这次的链路追踪
	Ignore()
	// IsIgnore 是否忽略
	IsIgnore() bool
	// GetLevel 获取层级
	GetLevel() int
	// SetHttpRequest 设置Http请求出入参
	SetHttpRequest(url string, reqHead map[string]any, rspHead map[string]string, requestBody string, responseBody string, statusCode int)
	// 设置得到的数据量
	SetRows(rows int)
}
