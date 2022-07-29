package linq

// 支持比较的集合
type linqFromC[C comparable] struct {
	// source array
	source []C
}

func FromC[C comparable](source []C) linqFromC[C] {
	return linqFromC[C]{
		source: source,
	}
}

// Where 对数据进行筛选
func (receiver linqFromC[C]) Where(fn func(item C) bool) linqFromC[C] {
	var lst []C
	for _, item := range receiver.source {
		if fn(item) {
			lst = append(lst, item)
		}
	}
	receiver.source = lst
	return receiver
}

// Contains 查找数组是否包含某元素
func (receiver linqFromC[C]) Contains(t C) bool {
	for _, item := range receiver.source {
		if item == t {
			return true
		}
	}
	return false
}

// Remove 移除指定值的元素
func (receiver linqFromC[C]) Remove(val C) []C {
	var lst []C
	for _, item := range receiver.source {
		if item != val {
			lst = append(lst, item)
		}
	}
	receiver.source = lst
	return lst
}

// Count 获取数量
func (receiver linqFromC[C]) Count() int {
	return len(receiver.source)
}
