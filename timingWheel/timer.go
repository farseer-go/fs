package timingWheel

import (
	"runtime"
	"sync/atomic"
	"time"
	_ "unsafe" // 必须导入以使用 go:linkname
)

//go:linkname procyield runtime.procyield
func procyield(cycles uint32)

type Timer struct {
	Id                int64 // 唯一ID
	C                 chan time.Time
	duration          time.Duration // 实际时长
	remainingDuration time.Duration // 在level层剩余时长
	PlanAt            time.Time     // 计划执行时间
	isPrecision       bool          // 是否使用高精度时间格
	isStop            int32         // 是否停止（atomic）
}

func (receiver *Timer) Stop() {
	atomic.StoreInt32(&receiver.isStop, 1)
}

func (receiver *Timer) isStopped() bool {
	return atomic.LoadInt32(&receiver.isStop) == 1
}

func (receiver *Timer) precision() {
	// 阶段1：大段休眠 (距离目标 > 2ms)
	// 使用 for 循环防止 time.Sleep 提前唤醒（虽然 Go 1.23+ 已极大优化，但这是安全做法）
	for {
		remaining := time.Until(receiver.PlanAt)
		if remaining <= 2*time.Millisecond {
			break
		}
		// 预留 2ms，避开操作系统调度不确定性
		time.Sleep(remaining - 2*time.Millisecond)
	}

	// 阶段2：让权等待 (100μs ~ 2ms)
	// 1000 并发下，Gosched 允许 P 执行其他 G，避免 CPU 核心过早被 100% 锁死
	for {
		remaining := time.Until(receiver.PlanAt)
		if remaining <= 100*time.Microsecond {
			break
		}
		runtime.Gosched()
	}

	// 阶段3：极致自旋 (最后 100μs)
	// 加入 procyield 优化，减少对 CPU 流水线和内存总线的过度占用
	for time.Now().Before(receiver.PlanAt) {
		// 每次空转 30 个周期左右，这比纯空循环更节能且减少核心干扰
		procyield(30)
	}

	// 到达目标时间，立即执行后续逻辑
}
