package data

import (
	"github.com/farseernet/farseer.go/configure"
	"log"
	"testing"
)

func TestTableSet_ToList(t *testing.T) {
	// 设置配置默认值，模拟配置文件
	configure.SetDefault("Database.test", "DataType=MySql,PoolMaxSize=50,PoolMinSize=1,ConnectionString=root:steden@123@tcp(mysql:3306)/test?charset=utf8&parseTime=True&loc=Local")
	context := NewContext[TestMysqlContext]("test")

	list := context.User.Select("Age").ToList()
	log.Println(list)
}
