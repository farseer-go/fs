package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"strings"
	"time"
)

type TableSet[Table any] struct {
	// 上下文
	DbContext *DbContext
	// 表名
	tableName string
	// 对应的表
	po  Table
	db  *gorm.DB
	err error
}

// NewTableSet 初始化表模型
func NewTableSet[Table any](dbContext *DbContext, tableName string, po Table) TableSet[Table] {
	return TableSet[Table]{
		po:        po,
		DbContext: dbContext,
		tableName: tableName,
	}
}

// 初始化Orm
func (table TableSet[Table]) data() *gorm.DB {
	if table.db == nil { // Data Source Name，参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name
		table.db, table.err = gorm.Open(table.getDriver(), &gorm.Config{})
		if table.err != nil {
			panic(table.err.Error())
		}
		table.db.Table(table.tableName)
		table.setPool()
	}
	return table.db
}

func (table TableSet[Table]) getDriver() gorm.Dialector {
	// 参考：https://gorm.cn/zh_CN/docs/connecting_to_the_database.html
	switch strings.ToLower(table.DbContext.dbConfig.DataType) {
	case "mysql":
		return mysql.Open(table.DbContext.dbConfig.ConnectionString)
	case "postgresql":
		return postgres.Open(table.DbContext.dbConfig.ConnectionString)
	case "sqlite":
		return sqlite.Open(table.DbContext.dbConfig.ConnectionString)
	case "sqlserver":
		return sqlserver.Open(table.DbContext.dbConfig.ConnectionString)
	}
	panic("无法识别数据库类型：" + table.DbContext.dbConfig.DataType)
}

func (table TableSet[Table]) setPool() {
	sqlDB, _ := table.db.DB()
	sqlDB.SetMaxIdleConns(table.DbContext.dbConfig.PoolMinSize) // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(table.DbContext.dbConfig.PoolMaxSize) // SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(time.Hour)                         // SetConnMaxLifetime 设置了连接可复用的最大时间。
}

func (table TableSet[Table]) Select(query interface{}, args ...interface{}) TableSet[Table] {
	table.data().Select(query, args)
	return table
}

func (table TableSet[Table]) Where(query interface{}, args ...interface{}) TableSet[Table] {
	table.data().Where(query)
	return table
}

func (table TableSet[Table]) Order(value interface{}) TableSet[Table] {
	table.data().Order(value)
	return table
}

func (table TableSet[Table]) ToList() []Table {
	var lst []Table
	table.data().Find(&lst)
	return lst
}

func (table TableSet[Table]) ToEntity() Table {
	var entity Table
	table.data().First(&entity)
	return entity
}

func (table TableSet[Table]) Count() int64 {
	var count int64
	table.data().Count(&count)
	return count
}

func (table TableSet[Table]) IsExists() bool {
	var count int64
	table.data().Count(&count)
	return count > 0
}

func (table TableSet[Table]) Insert(po Table) {
	table.data().Create(po)
}

func (table TableSet[Table]) Update(po Table) int64 {
	var count int64
	table.data().Updates(po)
	return count
}

func (table TableSet[Table]) UpdateValue(column string, value interface{}) {
	table.data().Update(column, value)
}

func (table TableSet[Table]) Delete() {
	table.data().Delete(nil)
}
