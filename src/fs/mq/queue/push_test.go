package queue

import (
	"fmt"
	"testing"
	"time"
)

func init() {
	Subscribe("test", 2, testQueue{})
}

func TestPush(t *testing.T) {
	for i := 0; i < 100; i++ {
		Push("test", i)
	}
	time.Sleep(time.Hour)
}

type testQueue struct{}

func (testQueue) Consumer(message []any) {
	fmt.Println(message)
}
