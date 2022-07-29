package linq

// 对集合进行排序
type linqFromOrder[T1 any, T2 Ordered] struct {
	// source array
	source *[]T1
}

// FromOrder 对集合进行排序
func FromOrder[T1 any, T2 Ordered](source []T1) linqFromOrder[T1, T2] {
	return linqFromOrder[T1, T2]{
		source: &source,
	}
}

// Where 对数据进行筛选
func (receiver linqFromOrder[T, T2]) Where(fn func(item T) bool) linqFromOrder[T, T2] {
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
func (receiver linqFromOrder[T, T2]) OrderBy(fn func(item T) T2) []T {
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
func (receiver linqFromOrder[T, T2]) OrderByDescending(fn func(item T) T2) []T {
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
func (receiver linqFromOrder[T, T2]) Min(fn func(item T) T2) T2 {
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
func (receiver linqFromOrder[T, T2]) Max(fn func(item T) T2) T2 {
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
