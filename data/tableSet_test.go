package data

import (
	"fmt"
	"github.com/farseernet/farseer.go/configure"
	"testing"
)

func TestTableSet_ToList(t *testing.T) {
	// 设置配置默认值，模拟配置文件
	configure.SetDefault("Database.test", "DataType=MySql,PoolMaxSize=50,PoolMinSize=1,ConnectionString=root:steden@123@tcp(mysql:3306)/test?charset=utf8&parseTime=True&loc=Local")
	context := Init[TestMysqlContext]("test")

	list := context.User.Select("Age").ToList()
	fmt.Println(list)
}
