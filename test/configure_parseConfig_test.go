package test

import (
	"github.com/farseer-go/fs/configure"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testConfig struct {
	dbName           string
	DataType         string
	PoolMaxSize      int
	PoolMinSize      int
	ConnectionString string
}

func TestParseConfig(t *testing.T) {
	conf := "dataType=MySql,poolMaxSize=50,poolMinSize=1,connectionString=root:steden@123@tcp(mysql:3306)/fops?charset=utf8&parseTime=True&loc=Local"
	dbConfig := configure.ParseConfig[testConfig](conf)
	assert.Equal(t, dbConfig.PoolMaxSize, 50)
	assert.Equal(t, dbConfig.PoolMinSize, 1)
	assert.Equal(t, dbConfig.DataType, "MySql")
}
