package test

import (
	"github.com/farseer-go/fs/snowflake"
	"runtime"
	"testing"
)

func TestSnowflake(t *testing.T) {
	runtime.GOMAXPROCS(1024)
	snowflake.Init(8, 12)
	for i := 0; i < 100000; i++ {
		snowflake.GenerateId()
	}
}
