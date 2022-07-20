package linq

type linqBy[T comparable] struct {
	// source array
	source []T
}

func By[T comparable](source []T) linqBy[T] {
	return linqBy[T]{
		source: source,
	}
}

// Where 对数据进行筛选
func (receiver linqBy[T]) Where(fn WhereFunc[T]) linqBy[T] {
	var lst []T
	for _, item := range receiver.source {
		if fn(item) {
			lst = append(lst, item)
		}
	}
	receiver.source = lst
	return receiver
}

// Contains 查找数组是否包含某元素
func (receiver linqBy[T]) Contains(t T) bool {
	for _, item := range receiver.source {
		if item == t {
			return true
		}
	}
	return false
}

// Remove 移除指定值的元素
func (receiver linqBy[T]) Remove(val T) []T {
	var lst []T
	for _, item := range receiver.source {
		if item != val {
			lst = append(lst, item)
		}
	}
	receiver.source = lst
	return lst
}

// Count 获取数量
func (receiver linqBy[T]) Count() int {
	return len(receiver.source)
}
