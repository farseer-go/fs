package queue

import (
	"log"
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
	time.Sleep(5 * time.Second)
}

func consumer(subscribeName string, message []any, remainingCount int) {
	log.Println("Name=", subscribeName, "，Msg=", message, "RemainingCount=", remainingCount)
}
