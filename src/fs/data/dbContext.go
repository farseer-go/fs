package data

import (
	"github.com/farseernet/farseer.go/configure"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

type DbContext struct {
	// 数据库配置
	dbConfig *dbConfig
}

// NewDbContext 初始化上下文
func NewDbContext(dbName string) *DbContext {
	configString := configure.GetString("Database." + dbName)
	if configString == "" {
		panic("[farseer.yaml]找不到相应的配置：Database.\" + dbName")
	}
	dbConfig := configure.ParseConfig[dbConfig](configString)
	dbContext := &DbContext{
		dbConfig: &dbConfig,
	}
	dbContext.dbConfig.dbName = dbName
	return dbContext
}

// Init 数据库上下文初始化
// dbName：数据库配置名称
func Init[TDbContext any](dbName string) *TDbContext {
	if dbName == "" {
		panic("dbName入参必须设置有效的值")
	}
	dbConfig := NewDbContext(dbName) // 嵌入类型
	//var dbName string       // 数据库配置名称
	customContext := new(TDbContext)
	contextValueOf := reflect.ValueOf(customContext).Elem()

	for i := 0; i < contextValueOf.NumField(); i++ {
		field := contextValueOf.Field(i)
		if !field.CanSet() {
			continue
		}
		data := contextValueOf.Type().Field(i).Tag.Get("data")
		var tableName string
		if strings.HasPrefix(data, "name=") {
			tableName = data[len("name="):]
		}
		if tableName == "" {
			panic("表名未设置，需要设置tag标签:data.name 的value")
		}
		// 再取tableSet的子属性，并设置值
		field.Addr().MethodByName("Init").Call([]reflect.Value{reflect.ValueOf(dbConfig), reflect.ValueOf(tableName)})
		//field.FieldByName("dbContext").Set(reflect.ValueOf(dbConfig))
		//field.FieldByName("tableName").Set(reflect.ValueOf(tableName))
	}

	/*	for i := 0; i < contextValueOf.NumField(); i++ {
		field := contextValueOf.Field(i)
		// 如果是上下文
		if field.Type() == reflect.TypeOf(&dbContext{}) {
			data := contextValueOf.Type().Field(i).Tag.Get("data")
			if !strings.HasPrefix(data, "name=") {
				panic("在" + field.Type().String() + "上下文中，必须为data.DbContext类型设置tag标签:data.name 的key")
			}
			dbName = data[len("name="):]
			if dbName == "" {
				panic("在" + field.Type().String() + "上下文中，必须为data.DbContext类型设置tag标签:data.name 的value")
			}
			// 设置上下文
			dbConfig = NewDbContext(dbName)
			field.Set(reflect.ValueOf(dbConfig))
		}
	}*/
	return customContext
}

// 获取对应驱动
func (dbContext *DbContext) getDriver() gorm.Dialector {
	// 参考：https://gorm.cn/zh_CN/docs/connecting_to_the_database.html
	switch strings.ToLower(dbContext.dbConfig.DataType) {
	case "mysql":
		return mysql.Open(dbContext.dbConfig.ConnectionString)
	case "postgresql":
		return postgres.Open(dbContext.dbConfig.ConnectionString)
	case "sqlite":
		return sqlite.Open(dbContext.dbConfig.ConnectionString)
	case "sqlserver":
		return sqlserver.Open(dbContext.dbConfig.ConnectionString)
	}
	panic("无法识别数据库类型：" + dbContext.dbConfig.DataType)
}
