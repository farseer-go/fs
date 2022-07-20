package parse

// Convert 通用的类型转换
func Convert[T any](source any, defVal T) T {
	switch source.(type) {
	case int:
		val := source.(int)
		var newVal T
		NumberToType(val, defVal, &newVal)
		return newVal
	}

	return defVal
}

// NumberToType 数字转换类型
func NumberToType[T numberType](source T, defVal any, newVal any) {
	switch defVal.(type) {
	case bool:
		isTrue := source == 1
		newVal = any(isTrue)
		return
	default:
		newVal = &defVal
	}

}

type numberType interface {
	~int | ~int64 | ~int8 | ~int16 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}
