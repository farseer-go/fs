package stopwatch

import (
	"strconv"
	"time"

	"github.com/farseer-go/fs/flog"
)

type Watch struct {
	// 执行起点时间
	startTime time.Time
	// 是否正在运行
	isRunning bool
	d         time.Duration
	// // 上次停止时已执行的时间点（毫秒）
	// lastElapsedMilliseconds int64
	// // 上次停止时已执行的时间点（微秒）
	// lastElapsedMicroseconds int64
	// // 上次停止时已执行的时间点（纳秒）
	// lastElapsedNanoseconds int64
}

// StartNew 创建计时器，并开始计时
func StartNew() *Watch {
	return &Watch{
		startTime: time.Now(),
		isRunning: true,
		d:         time.Duration(0),
	}
}

// New 创建计时器，并开始计时
func New() *Watch {
	return &Watch{
		startTime: time.Time{},
		isRunning: false,
		d:         time.Duration(0),
	}
}

// Restart 重置计时器
func (sw *Watch) Restart() {
	sw.startTime = time.Now()
	sw.isRunning = true
	sw.d = time.Duration(0)
}

// Start 继续计时
func (sw *Watch) Start() {
	sw.startTime = time.Now()
	sw.isRunning = true
}

// Stop 停止计时
func (sw *Watch) Stop() {
	sub := time.Since(sw.startTime)
	sw.d += sub
	sw.isRunning = false
}

// ElapsedDuration 返回当前已计时的时间
func (sw *Watch) ElapsedDuration() time.Duration {
	if sw.isRunning {
		return (time.Since(sw.startTime) + sw.d)
	}
	return sw.d
}

// GetMillisecondsText 返回当前已计时的时间（毫秒）
func (sw *Watch) GetMillisecondsText() string {
	return flog.Red(strconv.FormatInt(sw.ElapsedDuration().Milliseconds(), 10) + " ms")
}

// GetMicrosecondsText 返回当前已计时的时间（微秒）
func (sw *Watch) GetMicrosecondsText() string {
	return flog.Red(strconv.FormatInt(sw.ElapsedDuration().Microseconds(), 10) + " us")
}

// GetNanosecondsText 返回当前已计时的时间（纳秒）
func (sw *Watch) GetNanosecondsText() string {
	return flog.Red(strconv.FormatInt(sw.ElapsedDuration().Nanoseconds(), 10) + " ns")
}

// GetNanosecondsText 返回当前已计时的时间
func (sw *Watch) GetText() string {
	return flog.Red(sw.ElapsedDuration().String())
}
