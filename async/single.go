package async

import "sync"

// call 代表一个正在进行中或已完成的请求
type call struct {
	wg  sync.WaitGroup
	val any
	err error
}

// Single 管理具有相同 key 的并发请求，实现并发请求时，只有第一个请求会被执行，其他请求等待结果返回
// 示例用法：
//
//	var single async.Single
//	result, err := single.Do("key", func() (any, error) {
//		// 执行实际操作
//	})
type Single struct {
	mu sync.Mutex       // 保护 m
	m  map[string]*call // 懒加载，存放 key 对应的请求
}

// Do 执行函数，确保相同 key 的请求在同一时间只执行一次
func (receiver *Single) Do(key string, fn func() (any, error)) (any, error) {
	receiver.mu.Lock()
	if receiver.m == nil {
		receiver.m = make(map[string]*call)
	}

	// 如果 key 已存在，表示有人正在处理
	if c, ok := receiver.m[key]; ok {
		receiver.mu.Unlock()
		c.wg.Wait() // 坐下等现成的
		return c.val, c.err
	}

	// 我是第一个到的，这时创建一个 call 并放入 map，同时释放锁，这时允许其它协程进来排队
	c := new(call)
	c.wg.Add(1)
	receiver.m[key] = c
	receiver.mu.Unlock()

	// 使用 defer 确保即使 fn 抛出 panic，也不会导致其他协程永久阻塞
	defer func() {
		receiver.mu.Lock()
		delete(receiver.m, key)
		receiver.mu.Unlock()
		c.wg.Done() // 释放所有等待者
	}()

	// 必须赋值给 c，这样其他等待的 goroutine 才能看到结果
	c.val, c.err = fn()
	return c.val, c.err
}
