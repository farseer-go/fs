package linq

import (
	"github.com/farseernet/farseer.go/core"
	"github.com/farseernet/farseer.go/utils/parse"
	"reflect"
)

// 数据对集合数据筛选
type linqForm[T any] struct {
	// source array
	source []T
}

// From 数据对集合数据筛选
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

// First 查找符合条件的第一个元素
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

// RemoveItem 移除指定值的元素
func (receiver linqForm[T]) RemoveItem(t T) []T {
	var lst []T
	for _, item := range receiver.source {
		if reflect.ValueOf(item) != reflect.ValueOf(t) {
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

// Any 是否存在
func (receiver linqForm[T]) Any() bool {
	return len(receiver.source) > 0
}

// All 是否所有数据都满足fn条件
func (receiver linqForm[T]) All(fn func(item T) bool) bool {
	for _, item := range receiver.source {
		if !fn(item) {
			return false
		}
	}
	return true
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

// ContainsItem 查找数组是否包含某元素
func (receiver linqForm[T]) ContainsItem(t T) bool {
	for _, item := range receiver.source {
		if reflect.ValueOf(item) == reflect.ValueOf(t) {
			return true
		}
	}
	return false
}

// Select 筛选子元素字段
//
// arrSlice：切片数组类型
//
// eg:
// 	lstYaml := []string{"1", "", "2"}
// 	var lst []string
//
// 	From(lstYaml).Select(&lst, func(item string) any {
// 	    return "go:" + item
// 	})
//
// 	result:
// 	lst = []string { "go:1", "go:", "go:2" }
func (receiver linqForm[T]) Select(arrSlice any, fn func(item T) any) {
	arrVal := reflect.ValueOf(arrSlice).Elem()
	if arrVal.Kind() != reflect.Slice {
		panic("arr入参必须为切片类型")
	}

	// 定义反射类型的切片
	var lst = make([]reflect.Value, 0)
	for _, item := range receiver.source {
		lst = append(lst, reflect.ValueOf(fn(item)))
	}

	value := reflect.Append(arrVal, lst...)
	arrVal.Set(value)
}

// GroupBy 将数组进行分组后返回map
//
// eg:
// 	type testItem struct {
// 	  name string
// 	  age  int
//	}
//	lst := []testItem{{name: "steden", age: 36}, {name: "steden", age: 18}, {name: "steden2", age: 40}}
//	var lstMap map[string][]testItem
//	From(lst).GroupBy(&lstMap, func(item testItem) any {
// 	  return item.name
//	})
//
// 	result:
// 	lstMap = map[string][]testItem {
// 	  "steden": {
// 	    {name: "steden", age: 36},
// 	    {name: "steden", age: 18},
// 	  },
// 	  "steden2": {
// 	      {name: "steden2", age: 40},
// 	    },
// 	  }
func (receiver linqForm[T]) GroupBy(mapSlice any, getMapKeyFunc func(item T) any) {
	mapSliceVal := reflect.ValueOf(mapSlice).Elem()
	if mapSliceVal.Kind() != reflect.Map {
		panic("mapSlice入参必须为map类型")
	}

	mapSliceVal.Set(reflect.MakeMap(mapSliceVal.Type()))

	for _, item := range receiver.source {
		key := reflect.ValueOf(getMapKeyFunc(item))
		mapValue := mapSliceVal.MapIndex(key)
		if mapValue == reflect.ValueOf(nil) {
			mapValue = reflect.ValueOf([]T{item})
			mapSliceVal.SetMapIndex(key, mapValue)
		} else {
			mapValue = reflect.Append(mapValue, reflect.ValueOf(item))
			mapSliceVal.SetMapIndex(key, mapValue)
		}
	}
}

// OrderBy 正序排序，fn 返回的是要排序的字段的值
func (receiver linqForm[T]) OrderBy(fn func(item T) any) []T {
	lst := receiver.source

	// 首先拿数组第0个出来做为左边值
	for leftIndex := 0; leftIndex < len(lst); leftIndex++ {
		// 拿这个值与后面的值作比较
		leftValue := fn(lst[leftIndex])

		// 再拿出左边值索引后面的值一一对比
		for rightIndex := leftIndex + 1; rightIndex < len(lst); rightIndex++ {
			rightValue := fn(lst[rightIndex]) // 这个就是后面的值，会陆续跟数组后面的值做比较
			rightItem := lst[rightIndex]

			// 后面的值比前面的值小，说明要交换数据
			if compareLeftGreaterThanRight(leftValue, rightValue) {
				// 开始交换数据，先从后面交换到前面
				for swapIndex := rightIndex; swapIndex > leftIndex; swapIndex-- {
					lst[swapIndex] = lst[swapIndex-1]
				}
				lst[leftIndex] = rightItem
				leftValue = fn(lst[leftIndex])
			}
		}
	}
	return lst
}

// OrderByItem 正序排序，fn 返回的是要排序的字段的值
func (receiver linqForm[T]) OrderByItem() []T {
	lst := receiver.source

	// 首先拿数组第0个出来做为左边值
	for leftIndex := 0; leftIndex < len(lst); leftIndex++ {
		// 拿这个值与后面的值作比较
		leftValue := lst[leftIndex]

		// 再拿出左边值索引后面的值一一对比
		for rightIndex := leftIndex + 1; rightIndex < len(lst); rightIndex++ {
			rightValue := lst[rightIndex] // 这个就是后面的值，会陆续跟数组后面的值做比较
			rightItem := lst[rightIndex]

			// 后面的值比前面的值小，说明要交换数据
			if compareLeftGreaterThanRight(leftValue, rightValue) {
				// 开始交换数据，先从后面交换到前面
				for swapIndex := rightIndex; swapIndex > leftIndex; swapIndex-- {
					lst[swapIndex] = lst[swapIndex-1]
				}
				lst[leftIndex] = rightItem
				leftValue = lst[leftIndex]
			}
		}
	}
	return lst
}

// OrderByDescending 正序排序，倒序排序，fn 返回的是要排序的字段的值
func (receiver linqForm[T]) OrderByDescending(fn func(item T) any) []T {
	lst := receiver.source

	// 首先拿数组第0个出来做为左边值
	for leftIndex := 0; leftIndex < len(lst); leftIndex++ {
		// 拿这个值与后面的值作比较
		leftValue := fn(lst[leftIndex])

		// 再拿出左边值索引后面的值一一对比
		for rightIndex := leftIndex + 1; rightIndex < len(lst); rightIndex++ {
			rightValue := fn(lst[rightIndex]) // 这个就是后面的值，会陆续跟数组后面的值做比较
			rightItem := lst[rightIndex]

			// 后面的值比前面的值小，说明要交换数据
			if !compareLeftGreaterThanRight(leftValue, rightValue) {
				// 开始交换数据，先从后面交换到前面
				for swapIndex := rightIndex; swapIndex > leftIndex; swapIndex-- {
					lst[swapIndex] = lst[swapIndex-1]
				}
				lst[leftIndex] = rightItem
				leftValue = fn(lst[leftIndex])
			}
		}
	}
	return lst
}

// OrderByDescendingItem 正序排序，倒序排序，fn 返回的是要排序的字段的值
func (receiver linqForm[T]) OrderByDescendingItem() []T {
	lst := receiver.source

	// 首先拿数组第0个出来做为左边值
	for leftIndex := 0; leftIndex < len(lst); leftIndex++ {
		// 拿这个值与后面的值作比较
		leftValue := lst[leftIndex]

		// 再拿出左边值索引后面的值一一对比
		for rightIndex := leftIndex + 1; rightIndex < len(lst); rightIndex++ {
			rightValue := lst[rightIndex] // 这个就是后面的值，会陆续跟数组后面的值做比较
			rightItem := lst[rightIndex]

			// 后面的值比前面的值小，说明要交换数据
			if !compareLeftGreaterThanRight(leftValue, rightValue) {
				// 开始交换数据，先从后面交换到前面
				for swapIndex := rightIndex; swapIndex > leftIndex; swapIndex-- {
					lst[swapIndex] = lst[swapIndex-1]
				}
				lst[leftIndex] = rightItem
				leftValue = lst[leftIndex]
			}
		}
	}
	return lst
}

// Min 获取最小值
func (receiver linqForm[T]) Min(fn func(item T) any) any {
	lst := receiver.source

	minValue := fn(lst[0])
	for index := 1; index < len(lst); index++ {
		value := fn(lst[index])
		if compareLeftGreaterThanRight(minValue, value) {
			minValue = value
		}
	}
	return minValue
}

// MinItem 获取最小值
func (receiver linqForm[T]) MinItem() T {
	lst := receiver.source

	minValue := lst[0]
	for index := 1; index < len(lst); index++ {
		value := lst[index]
		if compareLeftGreaterThanRight(minValue, value) {
			minValue = value
		}
	}
	return minValue
}

// Max 获取最大值
func (receiver linqForm[T]) Max(fn func(item T) any) any {
	lst := receiver.source

	maxValue := fn(lst[0])
	for index := 1; index < len(lst); index++ {
		value := fn(lst[index])
		if compareLeftGreaterThanRight(value, maxValue) {
			maxValue = value
		}
	}
	return maxValue
}

// MaxItem 获取最大值
func (receiver linqForm[T]) MaxItem() T {
	lst := receiver.source

	maxValue := lst[0]
	for index := 1; index < len(lst); index++ {
		value := lst[index]
		if compareLeftGreaterThanRight(value, maxValue) {
			maxValue = value
		}
	}
	return maxValue
}

// Sum 求总和
func (receiver linqForm[T]) Sum(fn func(item T) any) any {
	lst := receiver.source
	var sum any
	for index := 0; index < len(lst); index++ {
		sum = addition(sum, fn(lst[index]))
	}
	return sum
}

// SumItem 求总和
func (receiver linqForm[T]) SumItem() T {
	lst := receiver.source
	var sum T
	for index := 0; index < len(lst); index++ {
		sum = addition(sum, lst[index]).(T)
	}
	return sum
}

// Avg 求平均数
func (receiver linqForm[T]) Avg(fn func(item T) any) float64 {
	sum := receiver.Sum(fn)
	count := len(receiver.source)
	return parse.Convert(sum, float64(0)) / parse.Convert(count, float64(0))
}

// AvgItem 求平均数
func (receiver linqForm[T]) AvgItem() float64 {
	sum := receiver.Sum(func(item T) any { return item })
	count := len(receiver.source)
	return parse.Convert(sum, float64(0)) / parse.Convert(count, float64(0))
}
