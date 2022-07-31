package linq

// 分组
type linqFormGroupBy[T1 any, T2 comparable] struct {
	// source array
	source []T1
}

// FormGroupBy 分组
func FormGroupBy[T1 any, T2 comparable](source []T1) linqFormGroupBy[T1, T2] {
	return linqFormGroupBy[T1, T2]{
		source: source,
	}
}

// GroupBy 将数组进行分组后返回map
func (receiver linqFormGroupBy[T1, TKey]) GroupBy(groupByFunc func(item T1) TKey) map[TKey][]T1 {
	lst := make(map[TKey][]T1)
	for _, item := range receiver.source {
		key := groupByFunc(item)
		lst[key] = append(lst[key], item)
	}
	return lst
}
