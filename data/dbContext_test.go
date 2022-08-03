package data

import (
	"github.com/farseernet/farseer.go/configure"
	"testing"
)

type TestMysqlContext struct {
	User TableSet[UserPO] `data:"name=user"`
}

type UserPO struct {
	Id int `gorm:"primaryKey"`
	// 用户名称
	Name string
	// 用户年龄
	Age int
}

func TestInit(t *testing.T) {
	// 设置配置默认值，模拟配置文件
	configure.SetDefault("Database.test", "DataType=MySql,PoolMaxSize=50,PoolMinSize=1,ConnectionString=root:steden@123@tcp(mysql:3306)/test?charset=utf8&parseTime=True&loc=Local")
	context := NewContext[TestMysqlContext]("test")
	if context.User.GetTableName() != "user" {
		t.Errorf("context.User.GetTableName() != \"user\"")
	}
}
