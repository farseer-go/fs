package parse

import "reflect"

const oneInt int = 1
const oneInt8 int8 = 1
const oneInt16 int16 = 1
const oneInt32 int32 = 1
const oneInt64 int64 = 1
const oneUint uint = 1
const oneUint8 uint8 = 1
const oneUint16 uint16 = 1
const oneUint32 uint32 = 1
const oneUint64 uint64 = 1
const oneFloat32 float32 = 1
const oneFloat64 float64 = 1

func equalTo1(source any, sourceKind reflect.Kind) bool {
	var result bool
	switch sourceKind {
	case reflect.Int8:
		result = source.(int8) == oneInt8
	case reflect.Int16:
		result = source.(int16) == oneInt16
	case reflect.Int32:
		result = source.(int32) == oneInt32
	case reflect.Int64:
		result = source.(int64) == oneInt64
	case reflect.Int:
		result = source.(int) == oneInt
	case reflect.Uint:
		result = source.(uint) == oneUint
	case reflect.Uint8:
		result = source.(uint8) == oneUint8
	case reflect.Uint16:
		result = source.(uint16) == oneUint16
	case reflect.Uint32:
		result = source.(uint32) == oneUint32
	case reflect.Uint64:
		result = source.(uint64) == oneUint64
	case reflect.Float32:
		result = source.(float32) == oneFloat32
	case reflect.Float64:
		result = source.(float64) == oneFloat64
	}
	return result
}
