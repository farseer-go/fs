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
type rabbitConfig struct {
	Server   serverConfig
	Exchange []exchangeConfig
}

type exchangeConfig struct {
	ExchangeName       string // 交换器名称
	RoutingKey         string // RoutingKey
	ExchangeType       string // 交换器类型
	UseConfirmModel    bool   // 是否需要ACK
	AutoCreateExchange bool   // 交换器不存在，是否自动创建
}

// rabbitConfig 配置项
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
	rabbitConfigs := configure.ParseConfigs[rabbitConfig]("Rabbit")
	assert.Equal(t, 2, len(rabbitConfigs))
	assert.Equal(t, "rabbit1:5672", rabbitConfigs[0].Server.Server)
	assert.Equal(t, "farseer", rabbitConfigs[0].Server.UserName)
	assert.Equal(t, "farseer", rabbitConfigs[0].Server.Password)
	assert.Equal(t, "Ex1", rabbitConfigs[0].Exchange[0].ExchangeName)
	assert.Equal(t, "", rabbitConfigs[0].Exchange[0].RoutingKey)
	assert.Equal(t, "fanout", rabbitConfigs[0].Exchange[0].ExchangeType)
	assert.Equal(t, false, rabbitConfigs[0].Exchange[0].UseConfirmModel)
	assert.Equal(t, true, rabbitConfigs[0].Exchange[0].AutoCreateExchange)
	assert.Equal(t, "Ex2", rabbitConfigs[0].Exchange[1].ExchangeName)
	assert.Equal(t, "", rabbitConfigs[0].Exchange[1].RoutingKey)
	assert.Equal(t, "fanout", rabbitConfigs[0].Exchange[1].ExchangeType)
	assert.Equal(t, false, rabbitConfigs[0].Exchange[1].UseConfirmModel)
	assert.Equal(t, true, rabbitConfigs[0].Exchange[1].AutoCreateExchange)

	assert.Equal(t, "rabbit2:5672", rabbitConfigs[1].Server.Server)
	assert.Equal(t, "farseer", rabbitConfigs[1].Server.UserName)
	assert.Equal(t, "farseer", rabbitConfigs[1].Server.Password)
	assert.Equal(t, "Ex3", rabbitConfigs[1].Exchange[0].ExchangeName)
	assert.Equal(t, "", rabbitConfigs[1].Exchange[0].RoutingKey)
	assert.Equal(t, "fanout", rabbitConfigs[1].Exchange[0].ExchangeType)
	assert.Equal(t, false, rabbitConfigs[1].Exchange[0].UseConfirmModel)
	assert.Equal(t, true, rabbitConfigs[1].Exchange[0].AutoCreateExchange)
	assert.Equal(t, "Ex4", rabbitConfigs[1].Exchange[1].ExchangeName)
	assert.Equal(t, "", rabbitConfigs[1].Exchange[1].RoutingKey)
	assert.Equal(t, "fanout", rabbitConfigs[1].Exchange[1].ExchangeType)
	assert.Equal(t, false, rabbitConfigs[1].Exchange[1].UseConfirmModel)
	assert.Equal(t, true, rabbitConfigs[1].Exchange[1].AutoCreateExchange)
}
