package timingWheel

import (
	"time"
)

type Timer struct {
	Id                int64 // 唯一ID
	C                 chan time.Time
	duration          time.Duration // 实际时长
	remainingDuration time.Duration // 在level层剩余时长
	PlanAt            time.Time     // 计划执行时间
	isPrecision       bool          // 是否使用高精度时间格
	isStop            bool          // 是否停止
}

func (receiver *Timer) Stop() {
	receiver.isStop = true
}

// 精确控制时间
func (receiver *Timer) precision() {
	// 留出5ms做最后精确控制
	milliseconds := receiver.PlanAt.Sub(time.Now()).Milliseconds() - 4
	if milliseconds > 0 {
		time.Sleep(time.Duration(milliseconds) * time.Millisecond)
		// 每次休眠1ms
		for receiver.PlanAt.Sub(time.Now()).Microseconds() >= 1100 {
			time.Sleep(1 * time.Millisecond)
		}
		//flog.Debugf("休眠时间(%d):+%s %d us", receiver.Id, receiver.PlanAt.Format("15:04:05.000"), receiver.PlanAt.Sub(time.Now()).Microseconds())
		// 每次休眠0.2ms
		for receiver.PlanAt.Sub(time.Now()).Microseconds() >= 55 {
			time.Sleep(50 * time.Microsecond)
		}
	}
}
