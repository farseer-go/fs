package stopwatch

import (
	"time"
)

type Stopwatch struct {
	// 执行起点时间
	startTime time.Time
	// 是否正在运行
	isRunning bool
	// 上次停止时已执行的时间点
	lastElapsedMilliseconds int64
}

// StartNew 创建计时器，并开始计时
func StartNew() *Stopwatch {
	return &Stopwatch{
		startTime:               time.Now(),
		isRunning:               true,
		lastElapsedMilliseconds: 0,
	}
}

// Restart 重置计时器
func (sw *Stopwatch) Restart() {
	sw.startTime = time.Now()
	sw.isRunning = true
	sw.lastElapsedMilliseconds = 0
}

// Start 继续计时
func (sw *Stopwatch) Start() {
	sw.startTime = time.Now()
	sw.isRunning = true
}

// Stop 停止计时
func (sw *Stopwatch) Stop() {
	sw.lastElapsedMilliseconds += time.Now().Sub(sw.startTime).Milliseconds()
	sw.isRunning = false
}

// ElapsedMilliseconds 返回当前已计时的时间（毫秒）
func (sw *Stopwatch) ElapsedMilliseconds() int64 {
	elapsedMilliseconds := sw.lastElapsedMilliseconds
	if sw.isRunning {
		elapsedMilliseconds += time.Now().Sub(sw.startTime).Milliseconds()
	}
	return elapsedMilliseconds
}
