package trace

type ITraceDetail interface {
	ToString() string
	GetTraceDetail() *BaseTraceDetail
	End(err error)
	SetSql(DbName string, tableName string, sql string)
	// Ignore 忽略这次的链路追踪
	Ignore()
	// IsIgnore 是否忽略
	IsIgnore() bool
	// GetLevel 获取层级
	GetLevel() int
}
