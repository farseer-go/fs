package trace

type TraceDetailDatabase struct {
	DbName             string `json:",omitempty"` // 数据库名
	DbTableName        string `json:",omitempty"` // 表名
	DbSql              string `json:",omitempty"` // SQL
	DbConnectionString string `json:",omitempty"` // 连接字符串
	DbRowsAffected     int64  `json:"omitempty"`  // 影响行数
}

func (receiver *TraceDetailDatabase) SetSql(connectionString string, DbName string, tableName string, sql string, rowsAffected int64) {
	receiver.DbConnectionString = connectionString
	receiver.DbName = DbName
	receiver.DbTableName = tableName
	receiver.DbSql = sql
	receiver.DbRowsAffected = rowsAffected
}
