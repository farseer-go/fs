package test

import (
	"fmt"
	"github.com/farseer-go/fs/flog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlog(t *testing.T) {
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
	flog.ErrorIfExists(fmt.Errorf("test error"))
	flog.ErrorIfExists(nil)
	flog.Critical("")
	flog.Criticalf("")
	flog.Print("")
	flog.Printf("")
	flog.Println("")
	flog.ComponentInfo("task", "")
	flog.ComponentInfof("task", "")

	assert.Equal(t, "aaa"+flog.Red("b")+"aaa", flog.ReplaceRed("aaabaaa", "b"))
	assert.Equal(t, "aaa"+flog.Blue("b")+"aaa", flog.ReplaceBlue("aaabaaa", "b"))
	assert.Equal(t, "aaa"+flog.Green("b")+"aaa", flog.ReplaceGreen("aaabaaa", "b"))
	assert.Equal(t, "aaa"+flog.Yellow("b")+"aaa", flog.ReplaceYellow("aaabaaa", "b"))

	assert.Equal(t, "aaa"+flog.Red("b")+flog.Red("c")+"aaa", flog.ReplaceReds("aaabcaaa", "b", "c"))
	assert.Equal(t, "aaa"+flog.Blue("b")+flog.Blue("c")+"aaa", flog.ReplaceBlues("aaabcaaa", "b", "c"))
	assert.Equal(t, "aaa"+flog.Green("b")+flog.Green("c")+"aaa", flog.ReplaceGreens("aaabcaaa", "b", "c"))
	assert.Equal(t, "aaa"+flog.Yellow("b")+flog.Yellow("c")+"aaa", flog.ReplaceYellows("aaabcaaa", "b", "c"))

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

	/*
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
	*/
}
