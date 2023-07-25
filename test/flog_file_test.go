package test

import (
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/flog"
	"testing"
	"time"
)

func TestFlogFile(t *testing.T) {
	fs.Initialize[testModule]("flog test")
	flog.Info("测试日志文件")
	time.Sleep(time.Hour)
}
