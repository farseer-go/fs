package test

import (
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/modules"
	"github.com/stretchr/testify/assert"
	fs2 "io/fs"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestFlogFile(t *testing.T) {
	_ = os.RemoveAll("./log/")
	fs.Initialize[modules.FarseerKernelModule]("flog test")
	flog.Info("测试日志文件")
	time.Sleep(time.Second * 2)

	fileCount := 0
	_ = filepath.WalkDir("./log/", func(filePath string, dirInfo fs2.DirEntry, err error) error {
		if "./log/" == filePath {
			return nil
		}
		if dirInfo.IsDir() {
			return fs2.SkipDir
		}
		fileCount++
		return nil
	})

	assert.Greater(t, fileCount, 0)
}
