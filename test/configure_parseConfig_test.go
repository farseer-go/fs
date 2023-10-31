package test

import (
	"github.com/farseer-go/fs/configure"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testConfig struct {
	DataType         string
	PoolMaxSize      int
	PoolMinSize      int
	ConnectionString string
	UsePool          bool
}
type testGetSubNodesConfig struct {
	Server serverConfig
	Test   []exchangeConfig
}

type exchangeConfig struct {
	TestName        string // 交换器名称
	RoutingKey      string // RoutingKey
	TestType        string // 交换器类型
	UseConfirmModel bool   // 是否需要ACK
	AutoCreateTest  bool   // 交换器不存在，是否自动创建
}

// testGetSubNodesConfig 配置项
type serverConfig struct {
	Server   string // 服务端地址
	UserName string // 用户名
	Password string // 密码
}

func TestParseString(t *testing.T) {
	conf := "dataType=MySql,poolMaxSize=50,poolMinSize=1,connectionString=root:steden@123@tcp(mysql:3306)/fops?charset=utf8&parseTime=True&loc=Local,UsePool=true"
	dbConfig := configure.ParseString[testConfig](conf)
	assert.Equal(t, dbConfig.PoolMaxSize, 50)
	assert.Equal(t, dbConfig.PoolMinSize, 1)
	assert.Equal(t, dbConfig.DataType, "MySql")
	assert.Equal(t, dbConfig.UsePool, true)

	assert.Panics(t, func() {
		configure.ParseString[testConfig]("dataType")
		configure.ParseString[string]("")
	})
}

func TestParseConfigs(t *testing.T) {
	configure.InitConfig()
	testConfigs := configure.ParseConfigs[testGetSubNodesConfig]("TestGetSubNodes")
	assert.Equal(t, 2, len(testConfigs))
	assert.Equal(t, "test:8888", testConfigs[0].Server.Server)
	assert.Equal(t, "farseer", testConfigs[0].Server.UserName)
	assert.Equal(t, "farseer", testConfigs[0].Server.Password)
	assert.Equal(t, "Ex1", testConfigs[0].Test[0].TestName)
	assert.Equal(t, "", testConfigs[0].Test[0].RoutingKey)
	assert.Equal(t, "fanout", testConfigs[0].Test[0].TestType)
	assert.Equal(t, false, testConfigs[0].Test[0].UseConfirmModel)
	assert.Equal(t, true, testConfigs[0].Test[0].AutoCreateTest)
	assert.Equal(t, "Ex2", testConfigs[0].Test[1].TestName)
	assert.Equal(t, "", testConfigs[0].Test[1].RoutingKey)
	assert.Equal(t, "fanout", testConfigs[0].Test[1].TestType)
	assert.Equal(t, false, testConfigs[0].Test[1].UseConfirmModel)
	assert.Equal(t, true, testConfigs[0].Test[1].AutoCreateTest)

	assert.Equal(t, "test2:8888", testConfigs[1].Server.Server)
	assert.Equal(t, "farseer", testConfigs[1].Server.UserName)
	assert.Equal(t, "farseer", testConfigs[1].Server.Password)
	assert.Equal(t, "Ex3", testConfigs[1].Test[0].TestName)
	assert.Equal(t, "", testConfigs[1].Test[0].RoutingKey)
	assert.Equal(t, "fanout", testConfigs[1].Test[0].TestType)
	assert.Equal(t, false, testConfigs[1].Test[0].UseConfirmModel)
	assert.Equal(t, true, testConfigs[1].Test[0].AutoCreateTest)
	assert.Equal(t, "Ex4", testConfigs[1].Test[1].TestName)
	assert.Equal(t, "", testConfigs[1].Test[1].RoutingKey)
	assert.Equal(t, "fanout", testConfigs[1].Test[1].TestType)
	assert.Equal(t, false, testConfigs[1].Test[1].UseConfirmModel)
	assert.Equal(t, true, testConfigs[1].Test[1].AutoCreateTest)
}
