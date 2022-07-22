package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func NewTableSet[Table any](dbContext *DbContext, tableName string, po Table) TableSet[Table] {
	return TableSet[Table]{
		po:        po,
		DbContext: dbContext,
		tableName: tableName,
	}
}

// 初始化Orm
func (table TableSet[Table]) data() *gorm.DB {
	if table.db == nil {
		table.db, table.err = gorm.Open(mysql.New(mysql.Config{
			DriverName: "my_mysql_driver",
			DSN:        "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local", // Data Source Name，参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name
		}), &gorm.Config{})
		if table.err != nil {
			panic(table.err.Error())
		}
		table.db.Table(table.tableName)
	}
	return table.db
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
