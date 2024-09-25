package async

import (
	"errors"
	"fmt"
	"github.com/timandy/routine"
	"sync"
)

// worker 异步结构体
type worker struct {
	wg  *sync.WaitGroup // sync.WaitGroup
	err error           // 返回错误
}

func New() *worker {
	return &worker{
		wg: &sync.WaitGroup{},
	}
}

// Add 添加异步执行的方法
func (ac *worker) Add(fn func()) {
	ac.wg.Add(1)
	routine.Go(func() {
		ac.executeFunc(fn)
	})
}

// Add 添加异步执行的方法
func (ac *worker) AddGO(fn func()) {
	ac.wg.Add(1)
	go func() {
		ac.executeFunc(fn)
	}()
}

func (ac *worker) executeFunc(fn func()) {
	defer func() {
		// 异常处理
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				ac.err = err.(error)
			default:
				ac.err = errors.New(fmt.Sprint(err))
			}
		}
		ac.wg.Done()
	}()
	fn()
}

// ContinueWith 当并行任务执行完后，以非阻塞方式执行callbacks
func (ac *worker) ContinueWith(callbacks ...func()) {
	// 使用异步等待，并执行callbacks
	routine.Go(func() {
		ac.wg.Wait()
		for _, callback := range callbacks {
			callback()
		}
	})
}

// Wait 阻塞等待执行完成
func (ac *worker) Wait() error {
	ac.wg.Wait()
	return ac.err
}
