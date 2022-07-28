package eventBus

import (
	"fmt"
	"testing"
	"time"
)

func init() {
	Subscribe("test_event_subscribe", Consumer)
	Subscribe("test_event_subscribe", Consumer)
}

var count int

type testEventPublish struct {
	Name string
}

func Consumer(message any, ea EventArgs) {
	count++
	event := message.(testEventPublish)
	fmt.Println("ID=", ea.Id, "message=", event, "count=", count)
}

func TestPublishEvent(t *testing.T) {
	PublishEvent("test_event_subscribe", testEventPublish{Name: "aaa"})
	fmt.Println("send aaa finished")
	PublishEventAsync("test_event_subscribe", testEventPublish{Name: "bbb"})
	fmt.Println("send bbb finished")
	time.Sleep(2 * time.Second)
}
