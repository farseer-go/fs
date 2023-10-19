package trace

type ITraceDetail interface {
	ToString() string
	GetTraceDetail() *BaseTraceDetail
	End(err error)
	SetSql(DbName string, tableName string, sql string)
}
