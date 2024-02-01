package fastReflect

//func MakeSlice(typMeta *TypeMeta, arr []any) any {
//	// 创建数组（耗时65ns）
//	slice := reflect.MakeSlice(typMeta.ReflectType, len(arr), len(arr))
//	slicePtr := slice.Pointer()
//	for i := 0; i < len(arr); i++ {
//		// 找到当前索引位置的内存地址。起始位置 + 每个元素占用的字节大小 ，得到第N个索引的内存起始位置
//		itemPtr := unsafe.Pointer(slicePtr + uintptr(i)*typMeta.ItemMeta.Size)
//		switch typMeta.ItemMeta.Kind {
//		case reflect.Int:
//			*(*int)(itemPtr) = arr[i].(int)
//		case reflect.String:
//			*(*string)(itemPtr) = arr[i].(string)
//		default:
//			println(itemPtr)
//		}
//	}
//	return slice.Interface()
//}
