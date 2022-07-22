package queue

import (
	"fmt"
	"testing"
	"time"
)

func init() {
	Subscribe("test", "A", 2, consumer)
	Subscribe("test", "B", 4, consumer)
}

func TestPush(t *testing.T) {
	for i := 0; i < 100; i++ {
		Push("test", i)
	}
	time.Sleep(time.Hour)
}

func consumer(subscribeName string, message []any, remainingCount int) {
	fmt.Println("Name=", subscribeName, "，Msg=", message, "RemainingCount=", remainingCount)
}
