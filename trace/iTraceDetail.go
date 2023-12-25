package trace

type ITraceDetail interface {
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
}
