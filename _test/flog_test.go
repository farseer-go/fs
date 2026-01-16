package test

import (
	"fmt"
	"testing"

	"github.com/farseer-go/fs/color"
	"github.com/farseer-go/fs/flog"
	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, "aaa"+color.Red("b")+"aaa", color.ReplaceRed("aaabaaa", "b"))
	assert.Equal(t, "aaa"+color.Blue("b")+"aaa", color.ReplaceBlue("aaabaaa", "b"))
	assert.Equal(t, "aaa"+color.Green("b")+"aaa", color.ReplaceGreen("aaabaaa", "b"))
	assert.Equal(t, "aaa"+color.Yellow("b")+"aaa", color.ReplaceYellow("aaabaaa", "b"))

	assert.Equal(t, "aaa"+color.Red("b")+color.Red("c")+"aaa", color.ReplaceReds("aaabcaaa", "b", "c"))
	assert.Equal(t, "aaa"+color.Blue("b")+color.Blue("c")+"aaa", color.ReplaceBlues("aaabcaaa", "b", "c"))
	assert.Equal(t, "aaa"+color.Green("b")+color.Green("c")+"aaa", color.ReplaceGreens("aaabcaaa", "b", "c"))
	assert.Equal(t, "aaa"+color.Yellow("b")+color.Yellow("c")+"aaa", color.ReplaceYellows("aaabcaaa", "b", "c"))

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
