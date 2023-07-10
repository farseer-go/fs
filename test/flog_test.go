package test

import (
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/flog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlog(t *testing.T) {
	configure.InitConfig()
	flog.Trace("")
	flog.Tracef("")
	flog.Debug("")
	flog.Debugf("")
	flog.Info("")
	flog.Infof("")
	flog.Warning("")
	flog.Warningf("")
	_ = flog.Error("")
	_ = flog.Errorf("")
	flog.Critical("")
	flog.Criticalf("")
	flog.Print("")
	flog.Printf("")
	flog.Println("")
	flog.ComponentInfo("task", "")
	flog.ComponentInfof("task", "")

	flog.Red("")
	flog.Blue("")
	flog.Green("")
	flog.Yellow("")
	assert.Panics(t, func() {
		flog.Panic("test error")
	})

	assert.Panics(t, func() {
		flog.Panicf("test error:%s", "content")
	})

	configure.SetDefault("Log.LogLevel", "Trace")
	flog.InitLog()
	flog.Print("aaa")

	configure.SetDefault("Log.LogLevel", "debug")
	flog.InitLog()
	flog.Trace("aaa")

	configure.SetDefault("Log.LogLevel", "Information")
	flog.InitLog()
	flog.Debug("aaa")

	configure.SetDefault("Log.LogLevel", "Warning")
	flog.InitLog()
	flog.Info("aaa")

	configure.SetDefault("Log.LogLevel", "Error")
	flog.InitLog()
	flog.Warning("aaa")

	configure.SetDefault("Log.LogLevel", "Critical")
	flog.InitLog()
	flog.Error("aaa")
}
