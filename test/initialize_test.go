package test

import (
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/modules"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"strings"
	"testing"
)

type testModule struct {
}

var lst []int

func (t testModule) DependsModule() []modules.FarseerModule {
	return nil
}

func (t testModule) PreInitialize() {
	lst = append(lst, 1)
}

func (t testModule) Initialize() {
	lst = append(lst, 2)
}

func (t testModule) PostInitialize() {
	lst = append(lst, 3)
	fs.AddInitCallback(func() {})
}

func (t testModule) Shutdown() {
}

func TestInitialize(t *testing.T) {
	fs.Initialize[testModule]("unit test")
	assert.Equal(t, 3, len(lst))
	assert.Equal(t, 1, lst[0])
	assert.Equal(t, 2, lst[1])
	assert.Equal(t, 3, lst[2])

	assert.Equal(t, 3, strings.Count(fs.AppIp, "."))
	assert.Equal(t, 18, len(strconv.FormatInt(fs.AppId, 10)))
	assert.Equal(t, "unit test", fs.AppName)
	assert.Equal(t, os.Getppid(), fs.ProcessId)

	hostName, _ := os.Hostname()
	assert.Equal(t, hostName, fs.HostName)
}
