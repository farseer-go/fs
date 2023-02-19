package timingWheel

import (
	"time"
)

type Timer struct {
	C           chan time.Time
	duration    time.Duration // 在level层的实际时长
	planAt      time.Time     // 计划执行时间
	isPrecision bool          // planAt是否需要精准
	isStop      bool          // 是否停止
}

func (receiver *Timer) Stop() {
	receiver.isStop = true
}

// 精确控制时间
func (receiver *Timer) precision() {
	milli := receiver.planAt.UnixMilli()
	// 留出5ms做最后精确控制
	milliseconds := receiver.planAt.Sub(time.Now()).Milliseconds() - 5
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)

	// 每次休眠1ms
	for milli > time.Now().UnixMilli()+1 {
		time.Sleep(1 * time.Millisecond)
	}

	// 每次休眠0.2ms
	for milli > time.Now().UnixMilli() {
		time.Sleep(200 * time.Microsecond)
	}
	receiver.C <- receiver.planAt
}
