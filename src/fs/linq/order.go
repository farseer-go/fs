package linq

import (
	"reflect"
)

type linqOrder[T1 any, T2 Ordered] struct {
	// source array
	source *[]T1
	// array type
	arrayType reflect.Type
	// element type
	elementType reflect.Type
	// value type
	value reflect.Value
}

func Order[T1 any, T2 Ordered](source []T1) linqOrder[T1, T2] {
	return linqOrder[T1, T2]{
		source: &source,
	}
}

type selectFunc[T any, T2 Ordered] func(item T) T2

// Where 对数据进行筛选
func (receiver linqOrder[T, T2]) Where(fn whereFunc[T]) linqOrder[T, T2] {
	var lst []T
	for _, item := range *receiver.source {
		if fn(item) {
			lst = append(lst, item)
		}
	}
	receiver.source = &lst
	return receiver
}

// OrderBy 正序排序，fn 返回的是要排序的字段的值
func (receiver linqOrder[T, T2]) OrderBy(fn selectFunc[T, T2]) []T {
	lst := *receiver.source

	// 首先拿数组第0个出来做为左边值
	for leftIndex := 0; leftIndex < len(lst); leftIndex++ {
		// 拿这个值与后面的值作比较
		leftValue := fn(lst[leftIndex])

		// 再拿出左边值索引后面的值一一对比
		for rightIndex := leftIndex + 1; rightIndex < len(lst); rightIndex++ {
			rightValue := fn(lst[rightIndex]) // 这个就是后面的值，会陆续跟数组后面的值做比较
			rightItem := lst[rightIndex]

			// 后面的值比前面的值小，说明要交换数据
			if rightValue < leftValue || rightValue == leftValue {

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

// OrderByDescending 倒序排序，fn 返回的是要排序的字段的值
func (receiver linqOrder[T, T2]) OrderByDescending(fn selectFunc[T, T2]) []T {
	lst := *receiver.source

	// 首先拿数组第0个出来做为左边值
	for leftIndex := 0; leftIndex < len(lst); leftIndex++ {
		// 拿这个值与后面的值作比较
		leftValue := fn(lst[leftIndex])

		// 再拿出左边值索引后面的值一一对比
		for rightIndex := leftIndex + 1; rightIndex < len(lst); rightIndex++ {
			rightValue := fn(lst[rightIndex]) // 这个就是后面的值，会陆续跟数组后面的值做比较
			rightItem := lst[rightIndex]

			// 后面的值比前面的值小，说明要交换数据
			if rightValue > leftValue || rightValue == leftValue {

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

// Min 获取最小值
func (receiver linqOrder[T, T2]) Min(fn selectFunc[T, T2]) T2 {
	lst := *receiver.source

	minValue := fn(lst[0])
	for index := 1; index < len(lst); index++ {
		value := fn(lst[index])
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}

// Max 获取最大值
func (receiver linqOrder[T, T2]) Max(fn selectFunc[T, T2]) T2 {
	lst := *receiver.source

	maxValue := fn(lst[0])
	for index := 1; index < len(lst); index++ {
		value := fn(lst[index])
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

/*
0	3
1	5
2	1
3	8
4	0
5	3
6	7
7	1
--------------------
0	1
1	3
2	5
3	8
4	0
5	3
6	7
7	1
--------------------
0	0
1	1
2	3
3	5
4	8
5	3
6	7
7	1*/

/*
0	3
1	5
2	1
3	8
4	0
5	3
6	7
7	1
--------------------
0	0
1	5
2	3
3	8
4	1
5	3
6	7
7	1
--------------------
0	3
1	5
2	1
3	8
4	0
5	3
6	7
7	1
*/
