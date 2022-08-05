package tasks

import (
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	Run("testRun", true, 1*time.Second, testRunFN)
	time.Sleep(5 * time.Second)
}

func testRunFN(context *TaskContext) {
	fmt.Println(time.Now())
}
