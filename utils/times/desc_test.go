package times

import (
	"testing"
	"time"
)

func TestGetSubDesc(t *testing.T) {
	ts1 := time.Date(2022, 8, 1, 20, 54, 12, 0, time.Local)
	ts2 := time.Date(2022, 8, 1, 19, 22, 12, 0, time.Local)

	desc := GetSubDesc(ts1, ts2)
	if desc != "1 小时 32 分" {
		t.Error()
	}
}
