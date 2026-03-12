package stopwatch

import (
	"strconv"
	"time"
	_ "unsafe"

	"github.com/farseer-go/fs/color"
)

// nanotime 直接读取 monotonic clock，跳过 wall clock 的采集开销
//
//go:linkname nanotime runtime.nanotime
func nanotime() int64

type Watch struct {
	// monotonic 起点时间（纳秒）
	startNano int64
	// 累计已用时间（纳秒）
	elapsed int64
	// 是否正在运行
	isRunning bool
}

// StartNew 创建计时器，并开始计时
func StartNew() *Watch {
	return &Watch{
		startNano: nanotime(),
		isRunning: true,
	}
}

// New 创建计时器，不开始计时
func New() *Watch {
	return &Watch{}
}

// Restart 重置计时器并重新开始
func (sw *Watch) Restart() {
	sw.elapsed = 0
	sw.startNano = nanotime()
	sw.isRunning = true
}

// Start 开始/继续计时（已在运行时忽略）
func (sw *Watch) Start() {
	if sw.isRunning {
		return
	}
	sw.startNano = nanotime()
	sw.isRunning = true
}

// Stop 停止计时（已停止时忽略）
func (sw *Watch) Stop() {
	if !sw.isRunning {
		return
	}
	sw.elapsed += nanotime() - sw.startNano
	sw.isRunning = false
}

// ElapsedDuration 返回当前已计时的时间
func (sw *Watch) ElapsedDuration() time.Duration {
	if sw.isRunning {
		return time.Duration(nanotime() - sw.startNano + sw.elapsed)
	}
	return time.Duration(sw.elapsed)
}

// GetMillisecondsText 返回当前已计时的时间（毫秒）
func (sw *Watch) GetMillisecondsText() string {
	return color.Red(strconv.FormatInt(sw.ElapsedDuration().Milliseconds(), 10) + " ms")
}

// GetMicrosecondsText 返回当前已计时的时间（微秒）
func (sw *Watch) GetMicrosecondsText() string {
	return color.Red(strconv.FormatInt(sw.ElapsedDuration().Microseconds(), 10) + " us")
}

// GetNanosecondsText 返回当前已计时的时间（纳秒）
func (sw *Watch) GetNanosecondsText() string {
	return color.Red(strconv.FormatInt(sw.ElapsedDuration().Nanoseconds(), 10) + " ns")
}

// GetText 返回当前已计时的时间（可读格式，精度保留2位小数）
func (sw *Watch) GetText() string {
	ns := sw.ElapsedDuration().Nanoseconds()
	var text string
	switch {
	case ns < int64(time.Microsecond):
		text = strconv.FormatInt(ns, 10) + " ns"
	case ns < int64(time.Millisecond):
		text = strconv.FormatFloat(float64(ns)/1e3, 'f', 2, 64) + " µs"
	case ns < int64(time.Second):
		text = strconv.FormatFloat(float64(ns)/1e6, 'f', 2, 64) + " ms"
	case ns < int64(time.Minute):
		text = strconv.FormatFloat(float64(ns)/1e9, 'f', 2, 64) + " s"
	default:
		text = time.Duration(ns).String()
	}
	return color.Red(text)
}
