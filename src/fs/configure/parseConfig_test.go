package configure

import (
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
	dbConfig := ParseConfig[testConfig](conf)
	if dbConfig.PoolMaxSize != 50 {
		t.Error()
	}
	if dbConfig.PoolMinSize != 1 {
		t.Error()
	}
	if dbConfig.DataType != "MySql" {
		t.Error()
	}
}
