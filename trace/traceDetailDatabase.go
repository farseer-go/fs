package trace

type TraceDetailDatabase struct {
	DbName             string // 数据库名
	DbTableName        string // 表名
	DbSql              string // SQL
	DbConnectionString string // 连接字符串
	DbRowsAffected     int64  // 影响行数
}

func (receiver *TraceDetailDatabase) SetSql(connectionString string, DbName string, tableName string, sql string, rowsAffected int64) {
	receiver.DbConnectionString = connectionString
	receiver.DbName = DbName
	receiver.DbTableName = tableName
	receiver.DbSql = sql
	receiver.DbRowsAffected = rowsAffected
}
