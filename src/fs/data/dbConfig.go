package data

// 数据库配置
type dbConfig struct {
	dbName           string
	DataType         string
	PoolMaxSize      int
	PoolMinSize      int
	ConnectionString string
}
