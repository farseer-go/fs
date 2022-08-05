package tasks

import (
	"github.com/farseernet/farseer.go/utils/stopwatch"
	"time"
)

// TaskContext 运行任务的上下文
type TaskContext struct {
	// 下一次运行的时间
	nextRunAt time.Time
	sw        *stopwatch.Watch
}

// SetNextTime 设置休眠时间
func (receiver *TaskContext) SetNextTime(nextTime time.Time) {
	if nextTime.UnixMicro() < time.Now().UnixMicro() {
		panic("nextTime时间，必须在当前时间之后")
	}
	receiver.nextRunAt = nextTime
}

// SetNextDuration 设置休眠时间
func (receiver *TaskContext) SetNextDuration(d time.Duration) {
	if d < 1 {
		panic("d参数，必须大于0")
	}
	receiver.nextRunAt = time.Now().Add(d)
}
