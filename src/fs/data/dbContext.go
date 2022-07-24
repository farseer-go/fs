package data

import (
	"fs/configure"
)

type DbContext struct {
	// 数据库配置
	dbConfig dbConfig
}

// NewDbContext 初始化上下文
func NewDbContext(dbName string) *DbContext {
	return &DbContext{
		dbConfig: configure.ParseConfig[dbConfig](configure.GetString("Database." + dbName)),
	}
}
