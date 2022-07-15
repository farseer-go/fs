package queue

import (
	"fmt"
	"testing"
	"time"
)

func init() {
	Subscribe("test", "A", 2, testQueue{})
	Subscribe("test", "B", 4, testQueue{})
}

func TestPush(t *testing.T) {
	for i := 0; i < 100; i++ {
		Push("test", i)
	}
	time.Sleep(time.Hour)
}

type testQueue struct{}

func (testQueue) Consumer(subscribeName string, message []any, remainingCount int) {
	fmt.Println("Name=", subscribeName, "，Msg=", message, "RemainingCount=", remainingCount)
}
