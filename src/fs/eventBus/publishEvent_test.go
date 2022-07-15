package eventBus

import (
	"fmt"
	"testing"
)

func init() {
	Subscribe("test_event_subscribe", testEventSubscribe{})
	Subscribe("test_event_subscribe", testEventSubscribe{})
}

type testEventSubscribe struct {
	Name string
}

func (rec testEventSubscribe) Consumer(message any, ea EventArgs) {
	fmt.Println("ID=", ea.Id, "message=", message)
}

func TestPublishEvent(t *testing.T) {
	PublishEvent("test_event_subscribe", "aaa")
	PublishEventAsync("test_event_subscribe", "bbb")
}
