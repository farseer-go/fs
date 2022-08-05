package tasks

import (
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	Run("testRun", true, 1*time.Second, func(context *TaskContext) {
		fmt.Println(time.Time{}.Year())
	})
	time.Sleep(5 * time.Second)
}
