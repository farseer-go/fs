package fs

import (
	"github.com/farseer-go/fs/modules"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
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
}

func (t testModule) Shutdown() {
}

func TestInitialize(t *testing.T) {
	Initialize[testModule]("unit test")
	assert.Equal(t, 3, len(lst))
	assert.Equal(t, 1, lst[0])
	assert.Equal(t, 2, lst[1])
	assert.Equal(t, 3, lst[2])

	assert.Equal(t, 3, strings.Count(AppIp, "."))
	assert.Equal(t, 18, len(strconv.FormatInt(AppId, 10)))
	assert.Equal(t, "unit test", AppName)
	assert.Equal(t, os.Getppid(), ProcessId)
	assert.True(t, time.Now().Sub(StartupAt) <= 1*time.Millisecond)

	hostName, _ := os.Hostname()
	assert.Equal(t, hostName, HostName)
}
