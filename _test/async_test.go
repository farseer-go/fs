package test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/farseer-go/fs/async"
	"github.com/stretchr/testify/assert"
)

func TestAsync_ContinueWith(t *testing.T) {
	var lock sync.Mutex
	var count = 0
	lock.Lock()
	worker := async.New()
	worker.Add(func() {
		lock.Lock()
		defer lock.Unlock()
		count += 1
	})
	worker.Add(func() {
		lock.Lock()
		defer lock.Unlock()
		count += 2
	})
	worker.Add(func() {
		lock.Lock()
		defer lock.Unlock()
		count += 3
	})
	worker.Add(func() {
		lock.Lock()
		defer lock.Unlock()
		count += 4
	})
	worker.ContinueWith(func() {
		lock.Lock()
		defer lock.Unlock()
		count += 5
	})

	count = 10
	lock.Unlock()
	time.Sleep(10 * time.Millisecond)

	lock.Lock()
	defer lock.Unlock()
	assert.Equal(t, 25, count)
}

func TestAsync_Wait(t *testing.T) {
	var lock sync.Mutex
	var count = 0
	worker := async.New()
	worker.Add(func() {
		lock.Lock()
		defer lock.Unlock()
		count += 1
	})
	worker.Add(func() {
		lock.Lock()
		defer lock.Unlock()
		count += 2
	})
	worker.Add(func() {
		lock.Lock()
		defer lock.Unlock()
		count += 3
	})
	worker.Add(func() {
		lock.Lock()
		defer lock.Unlock()
		count += 4
	})
	worker.Wait()
	count *= 2
	assert.Equal(t, 20, count)
}

func TestAsync_Error(t *testing.T) {
	var count = 0
	var num = 0
	worker := async.New()
	worker.Add(func() {
		count = count / num
	})
	err := worker.Wait()
	assert.NotEqual(t, err, nil)

	worker = async.New()
	worker.Add(func() {
		panic("error")
	})
	err = worker.Wait()

	assert.NotEqual(t, err, nil)
}

func TestSingleDo(t *testing.T) {
	s := &async.Single{}
	var count int32
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Do("constant-key", func() (any, error) {
				atomic.AddInt32(&count, 1) // 理论上只会执行 1 次
				time.Sleep(100 * time.Millisecond)
				return "done", nil
			})
		}()
	}
	wg.Wait()

	if atomic.LoadInt32(&count) != 1 {
		t.Errorf("期望执行 1 次，实际执行了 %d 次", count)
	}
}
