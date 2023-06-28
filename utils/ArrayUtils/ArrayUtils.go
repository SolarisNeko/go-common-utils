package ArrayUtils

import "reflect"

// Contains 检查切片中是否包含指定的元素
func Contains(array interface{}, targetElement interface{}) bool {
	sliceValue := reflect.ValueOf(array)
	if sliceValue.Kind() != reflect.Slice {
		panic("Slice parameter is required")
	}

	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i).Interface()
		if reflect.DeepEqual(item, targetElement) {
			return true
		}
	}

	return false
}
