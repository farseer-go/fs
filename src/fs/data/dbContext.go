package data

import (
	"errors"
	"fmt"
	"fs/configure"
	"reflect"
	"strings"
)

type DbContext struct {
	// 数据库配置
	dbConfig dbConfig
}

// NewDbContext 初始化上下文
func NewDbContext(dbName string) *DbContext {
	config := configure.GetString("Database." + dbName)
	if config == "" {
		fmt.Println(errors.New("err：[farseer.yaml]找不到相应的配置：Database." + dbName))
	}
	return &DbContext{
		dbConfig: configure.ParseConfig[dbConfig](configure.GetString(config)),
	}
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
		field.MethodByName("SetTableName").Call([]reflect.Value{reflect.ValueOf(tableName)})
		field.FieldByName("DbContext").Set(reflect.ValueOf(dbConfig))
	}

	/*	for i := 0; i < contextValueOf.NumField(); i++ {
		field := contextValueOf.Field(i)
		// 如果是上下文
		if field.Type() == reflect.TypeOf(&DbContext{}) {
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
