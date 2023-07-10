package test

import (
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/flog2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlog(t *testing.T) {
	configure.InitConfig()
	flog2.Trace("")
	flog2.Tracef("")
	flog2.Debug("")
	flog2.Debugf("")
	flog2.Info("")
	flog2.Infof("")
	flog2.Warning("")
	flog2.Warningf("")
	_ = flog2.Error("")
	_ = flog2.Errorf("")
	flog2.Critical("")
	flog2.Criticalf("")
	flog2.Print("")
	flog2.Printf("")
	flog2.Println("")
	flog2.ComponentInfo("task", "")
	flog2.ComponentInfof("task", "")

	flog.Red("")
	flog.Blue("")
	flog.Green("")
	flog.Yellow("")
	assert.Panics(t, func() {
		flog2.Panic("test error")
	})

	assert.Panics(t, func() {
		flog2.Panicf("test error:%s", "content")
	})

	configure.SetDefault("Log.LogLevel", "Trace")
	flog.InitLog()
	flog2.Print("aaa")

	configure.SetDefault("Log.LogLevel", "debug")
	flog.InitLog()
	flog2.Trace("aaa")

	configure.SetDefault("Log.LogLevel", "Information")
	flog.InitLog()
	flog2.Debug("aaa")

	configure.SetDefault("Log.LogLevel", "Warning")
	flog.InitLog()
	flog2.Info("aaa")

	configure.SetDefault("Log.LogLevel", "Error")
	flog.InitLog()
	flog2.Warning("aaa")

	configure.SetDefault("Log.LogLevel", "Critical")
	flog.InitLog()
	flog2.Error("aaa")
}
