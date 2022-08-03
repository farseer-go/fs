package data

import (
	"github.com/farseernet/farseer.go/core"
	"gorm.io/gorm"
	"time"
)

// TableSet 数据库表操作
type TableSet[Table any] struct {
	// 上下文（用指针的方式，共享同一个上下文）
	dbContext *DbContext
	// 表名
	tableName string
	db        *gorm.DB
	err       error
}

// Init 在反射的时候会调用此方法
func (table *TableSet[Table]) Init(dbContext *DbContext, tableName string) {
	table.dbContext = dbContext
	table.SetTableName(tableName)
}

// SetTableName 设置表名
func (table *TableSet[Table]) SetTableName(tableName string) {
	table.tableName = tableName
	if table.db == nil {
		return
	}
	table.db.Table(table.tableName)
}

// GetTableName 获取表名称
func (table *TableSet[Table]) GetTableName() string {
	return table.tableName
}

// 初始化Orm
func (table *TableSet[Table]) data() *gorm.DB {
	if table.db == nil { // Data Source Name，参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name
		table.db, table.err = gorm.Open(table.dbContext.getDriver(), &gorm.Config{})
		if table.err != nil {
			panic(table.err.Error())
		}
		table.db.Table(table.tableName)
		table.setPool()
	}
	return table.db
}

// 设置池大小
func (table *TableSet[Table]) setPool() {
	sqlDB, _ := table.db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	if table.dbContext.dbConfig.PoolMinSize > 0 {
		sqlDB.SetMaxIdleConns(table.dbContext.dbConfig.PoolMinSize)
	}
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	if table.dbContext.dbConfig.PoolMaxSize > 0 {
		sqlDB.SetMaxOpenConns(table.dbContext.dbConfig.PoolMaxSize)
	}
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

// Select 筛选字段
func (table *TableSet[Table]) Select(query interface{}, args ...interface{}) *TableSet[Table] {
	table.data().Select(query, args...)
	return table
}

// Where 条件
func (table *TableSet[Table]) Where(query interface{}, args ...interface{}) *TableSet[Table] {
	table.data().Where(query, args...)
	return table
}

// Order 排序
func (table *TableSet[Table]) Order(value interface{}) *TableSet[Table] {
	table.data().Order(value)
	return table
}

// Desc 倒序
func (table *TableSet[Table]) Desc(fieldName string) *TableSet[Table] {
	table.data().Order(fieldName + " desc")
	return table
}

// Asc 正序
func (table *TableSet[Table]) Asc(fieldName string) *TableSet[Table] {
	table.data().Order(fieldName + " asc")
	return table
}

// Limit 限制记录数
func (table *TableSet[Table]) Limit(limit int) *TableSet[Table] {
	table.data().Limit(limit)
	return table
}

// ToList 返回结果集
func (table *TableSet[Table]) ToList() []Table {
	var lst []Table
	table.data().Find(&lst)
	return lst
}

// ToPageList 返回分页结果集
func (table *TableSet[Table]) ToPageList(pageSize int, pageIndex int) core.PageList[Table] {
	offset := (pageIndex - 1) * pageSize
	var lst []Table
	table.data().Offset(offset).Limit(pageSize).Find(&lst)

	return core.NewPageList[Table](lst, table.Count())
}

// ToEntity 返回单个对象
func (table *TableSet[Table]) ToEntity() Table {
	var entity Table
	table.data().Take(&entity)
	return entity
}

// Count 返回表中的数量
func (table *TableSet[Table]) Count() int64 {
	var count int64
	table.data().Count(&count)
	return count
}

// IsExists 是否存在记录
func (table *TableSet[Table]) IsExists() bool {
	var count int64
	table.data().Count(&count)
	return count > 0
}

// Insert 新增记录
func (table *TableSet[Table]) Insert(po *Table) {
	table.data().Create(po)
}

// Update 修改记录
func (table *TableSet[Table]) Update(po Table) int64 {
	result := table.data().Updates(po)
	return result.RowsAffected
}

// UpdateValue 修改单个字段
func (table *TableSet[Table]) UpdateValue(column string, value interface{}) {
	table.data().Update(column, value)
}

// Delete 删除记录
func (table *TableSet[Table]) Delete() int64 {
	result := table.data().Delete(nil)
	return result.RowsAffected
}

// GetString 获取单条记录中的单个string类型字段值
func (table *TableSet[Table]) GetString(fieldName string) string {
	rows, _ := table.data().Select(fieldName).Limit(1).Rows()
	defer rows.Close()
	var val string
	for rows.Next() {
		rows.Scan(&val)
		// ScanRows 方法用于将一行记录扫描至结构体
		//table.data().ScanRows(rows, &user)
	}
	return val
}

// GetInt 获取单条记录中的单个int类型字段值
func (table *TableSet[Table]) GetInt(fieldName string) int {
	rows, _ := table.data().Select(fieldName).Limit(1).Rows()
	defer rows.Close()
	var val int
	for rows.Next() {
		rows.Scan(&val)
		// ScanRows 方法用于将一行记录扫描至结构体
		//table.data().ScanRows(rows, &user)
	}
	return val
}

// GetLong 获取单条记录中的单个int64类型字段值
func (table *TableSet[Table]) GetLong(fieldName string) int64 {
	rows, _ := table.data().Select(fieldName).Limit(1).Rows()
	defer rows.Close()
	var val int64
	for rows.Next() {
		rows.Scan(&val)
		// ScanRows 方法用于将一行记录扫描至结构体
		//table.data().ScanRows(rows, &user)
	}
	return val
}

// GetBool 获取单条记录中的单个bool类型字段值
func (table *TableSet[Table]) GetBool(fieldName string) bool {
	rows, _ := table.data().Select(fieldName).Limit(1).Rows()
	defer rows.Close()
	var val bool
	for rows.Next() {
		rows.Scan(&val)
		// ScanRows 方法用于将一行记录扫描至结构体
		//table.data().ScanRows(rows, &user)
	}
	return val
}

// GetFloat32 获取单条记录中的单个float32类型字段值
func (table *TableSet[Table]) GetFloat32(fieldName string) float32 {
	rows, _ := table.data().Select(fieldName).Limit(1).Rows()
	defer rows.Close()
	var val float32
	for rows.Next() {
		rows.Scan(&val)
		// ScanRows 方法用于将一行记录扫描至结构体
		//table.data().ScanRows(rows, &user)
	}
	return val
}

// GetFloat64 获取单条记录中的单个float64类型字段值
func (table *TableSet[Table]) GetFloat64(fieldName string) float64 {
	rows, _ := table.data().Select(fieldName).Limit(1).Rows()
	defer rows.Close()
	var val float64
	for rows.Next() {
		rows.Scan(&val)
		// ScanRows 方法用于将一行记录扫描至结构体
		//table.data().ScanRows(rows, &user)
	}
	return val
}
