package test

import (
	"github.com/farseer-go/fs/async"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
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
