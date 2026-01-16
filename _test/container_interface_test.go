package test

type IDatabase interface {
	SetTableName(tableName string)
	GetTableName() string
	GetDbType() string
}

type mysql struct {
	tableName string
}

func (r *mysql) SetTableName(tableName string) { r.tableName = tableName }
func (r *mysql) GetTableName() string          { return r.tableName }
func (r *mysql) GetDbType() string             { return "mysql" }

type sqlserver struct {
	tableName string
}

func (r *sqlserver) SetTableName(tableName string) { r.tableName = tableName }
func (r *sqlserver) GetTableName() string          { return r.tableName }
func (r *sqlserver) GetDbType() string             { return "sqlserver" }
