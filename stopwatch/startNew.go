package stopwatch

import (
	"github.com/farseer-go/fs/flog"
	"strconv"
	"time"
)

type Watch struct {
	// 执行起点时间
	startTime time.Time
	// 是否正在运行
	isRunning bool
	// 上次停止时已执行的时间点（毫秒）
	lastElapsedMilliseconds int64
	// 上次停止时已执行的时间点（微秒）
	lastElapsedMicroseconds int64
	// 上次停止时已执行的时间点（纳秒）
	lastElapsedNanoseconds int64
}

// StartNew 创建计时器，并开始计时
func StartNew() *Watch {
	return &Watch{
		startTime:               time.Now(),
		isRunning:               true,
		lastElapsedMilliseconds: 0,
		lastElapsedMicroseconds: 0,
		lastElapsedNanoseconds:  0,
	}
}

// New 创建计时器，并开始计时
func New() *Watch {
	return &Watch{
		startTime:               time.Time{},
		isRunning:               false,
		lastElapsedMilliseconds: 0,
		lastElapsedMicroseconds: 0,
		lastElapsedNanoseconds:  0,
	}
}

// Restart 重置计时器
func (sw *Watch) Restart() {
	sw.startTime = time.Now()
	sw.isRunning = true
	sw.lastElapsedMilliseconds = 0
	sw.lastElapsedMicroseconds = 0
	sw.lastElapsedNanoseconds = 0
}

// Start 继续计时
func (sw *Watch) Start() {
	sw.startTime = time.Now()
	sw.isRunning = true
}

// Stop 停止计时
func (sw *Watch) Stop() {
	sub := time.Since(sw.startTime)
	sw.lastElapsedMilliseconds += sub.Milliseconds()
	sw.lastElapsedMicroseconds += sub.Microseconds()
	sw.lastElapsedNanoseconds += sub.Nanoseconds()
	sw.isRunning = false
}

// ElapsedMilliseconds 返回当前已计时的时间（毫秒）
func (sw *Watch) ElapsedMilliseconds() int64 {
	if sw.isRunning {
		return time.Since(sw.startTime).Milliseconds() + sw.lastElapsedMilliseconds
	}
	return sw.lastElapsedMilliseconds
}

// ElapsedMicroseconds 返回当前已计时的时间（微秒）
func (sw *Watch) ElapsedMicroseconds() int64 {
	if sw.isRunning {
		return time.Now().Sub(sw.startTime).Microseconds() + sw.lastElapsedMicroseconds
	}
	return sw.lastElapsedMicroseconds
}

// ElapsedNanoseconds 返回当前已计时的时间（纳秒）
func (sw *Watch) ElapsedNanoseconds() int64 {
	if sw.isRunning {
		return time.Since(sw.startTime).Nanoseconds() + sw.lastElapsedNanoseconds
	}
	return sw.lastElapsedNanoseconds
}

// GetMillisecondsText 返回当前已计时的时间（毫秒）
func (sw *Watch) GetMillisecondsText() string {
	return flog.Red(strconv.FormatInt(sw.ElapsedMilliseconds(), 10) + " ms")
}

// GetMicrosecondsText 返回当前已计时的时间（微秒）
func (sw *Watch) GetMicrosecondsText() string {
	return flog.Red(strconv.FormatInt(sw.ElapsedMicroseconds(), 10) + " us")
}

// GetNanosecondsText 返回当前已计时的时间（纳秒）
func (sw *Watch) GetNanosecondsText() string {
	return flog.Red(strconv.FormatInt(sw.ElapsedNanoseconds(), 10) + " ns")
}
