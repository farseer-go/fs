package tasks

import (
	"github.com/farseernet/farseer.go/utils/stopwatch"
	"log"
	"time"
)

// Run 运行一个任务
// isRunNow:是否立即执行fn任务
// interval:任务运行的间隔时间
// taskFn:要运行的任务
func Run(taskName string, isRunNow bool, interval time.Duration, taskFn func(context *TaskContext)) {
	// 不立即运行，则先休眠interval时间
	if interval <= 0 {
		panic("interval参数，必须大于0")
	}

	if !isRunNow {
		time.Sleep(interval)
	}

	go func() {
		for {
			runTask(taskName, interval, taskFn)
		}
	}()
}

// 运行任务
func runTask(taskName string, interval time.Duration, taskFn func(context *TaskContext)) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("taskFn [%s] throw exception：%s", taskName, r)
		}
	}()
	context := &TaskContext{
		sw: stopwatch.StartNew(),
	}
	taskFn(context)
	if context.nextRunAt.Year() >= 2022 {
		time.Sleep(context.nextRunAt.Sub(time.Now()))
	} else {
		time.Sleep(interval)
	}
}
