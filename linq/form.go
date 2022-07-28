package linq

import "github.com/farseernet/farseer.go/core"

type linqForm[T any] struct {
	// source array
	source []T
}

func From[T any](source []T) linqForm[T] {
	return linqForm[T]{
		source: source,
	}
}

// Where 对数据进行筛选
func (receiver linqForm[T]) Where(fn func(item T) bool) linqForm[T] {
	var lst []T
	for _, item := range receiver.source {
		if fn(item) {
			lst = append(lst, item)
		}
	}
	receiver.source = lst
	return receiver
}

// Find 查找符合条件的元素
func (receiver linqForm[T]) Find(fn func(item T) bool) T {
	for _, item := range receiver.source {
		if fn(item) {
			return item
		}
	}
	var t T
	return t
}

// FindAll 查找符合条件的元素列表
func (receiver linqForm[T]) FindAll(fn func(item T) bool) []T {
	var lst []T
	for _, item := range receiver.source {
		if fn(item) {
			lst = append(lst, item)
		}
	}
	return lst
}

// First 查找符合条件的元素
func (receiver linqForm[T]) First() T {
	if len(receiver.source) > 0 {
		return (receiver.source)[0]
	}
	var t T
	return t
}

// ToArray 查找符合条件的元素列表
func (receiver linqForm[T]) ToArray() []T {
	return receiver.source
}

// RemoveAll 移除条件=true的元素
func (receiver linqForm[T]) RemoveAll(fn func(item T) bool) []T {
	var lst []T
	for _, item := range receiver.source {
		if !fn(item) {
			lst = append(lst, item)
		}
	}
	receiver.source = lst
	return lst
}

// Count 获取数量
func (receiver linqForm[T]) Count() int {
	return len(receiver.source)
}

// ToPageList 数组分页
func (receiver linqForm[T]) ToPageList(pageSize int, pageIndex int) core.PageList[T] {
	pageList := core.PageList[T]{
		RecordCount: int64(len(receiver.source)),
	}

	if pageSize < 1 {
		pageSize = 10
	}

	// 计算总页数
	var allCurrentPage int64 = 0
	// 没有设置pageIndex，则按take返回
	if pageIndex < 1 {
		pageList.List = receiver.Take(pageSize)
		return pageList
	}

	allCurrentPage = pageList.RecordCount / int64(pageSize)
	if pageList.RecordCount%int64(pageSize) != 0 {
		allCurrentPage++
	}
	if allCurrentPage == 0 {
		allCurrentPage = 1
	}

	if int64(pageIndex) > allCurrentPage {
		pageIndex = int(allCurrentPage)
	}
	skipCount := pageSize * (pageIndex - 1)
	pageList.List = receiver.source[skipCount : skipCount+pageSize]
	return pageList
}

// Take 返回前多少条数据
func (receiver linqForm[T]) Take(count int) []T {
	recordCount := len(receiver.source)
	// 总长度比count小，则直接返回全部数据
	if recordCount < count {
		return receiver.source
	}

	return receiver.source[0:count]
}
