package test

import (
	"github.com/farseer-go/fs/configure"
	"github.com/farseer-go/fs/flog"
	"testing"
)

func TestFlog(t *testing.T) {

	configure.ReadInConfig()
	flog.Trace("")
	flog.Tracef("")
	flog.Debug("")
	flog.Debugf("")
	flog.Info("")
	flog.Infof("")
	flog.Warning("")
	flog.Warningf("")
	flog.Error("")
	flog.Errorf("")
	flog.Critical("")
	flog.Criticalf("")
	flog.Print("")
	flog.Printf("")
	flog.Println("")
	flog.ComponentInfo("task", "")
	flog.ComponentInfof("task", "")
}
