package timingWheel

import (
	"time"
)

// defaultTimingWheel 在farseer.yaml配置时间轮大小、格数
var defaultTimingWheel *timingWheel

func NewDefault(interval time.Duration, bucketsNum int) {
	defaultTimingWheel = New(interval, bucketsNum)
}

// Start 开始运行
func Start() {
	defaultTimingWheel.Start()
}

// Add 添加一个定时任务
func Add(d time.Duration) *Timer {
	return defaultTimingWheel.Add(d)
}

// AddTime 添加一个定时任务
func AddTime(t time.Time) *Timer {
	return defaultTimingWheel.AddTime(t)
}

// AddPrecision 添加一个定时任务（高精度）
func AddPrecision(d time.Duration) *Timer {
	return defaultTimingWheel.AddPrecision(d)
}

// AddTimePrecision 添加一个定时任务（高精度）
func AddTimePrecision(t time.Time) *Timer {
	return defaultTimingWheel.AddTimePrecision(t)
}
