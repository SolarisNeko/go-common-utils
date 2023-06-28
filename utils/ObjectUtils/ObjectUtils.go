package ObjectUtils

import "reflect"

// IsEquals 判断两个值是否相等
func IsEquals(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

func IsNotEquals(a, b interface{}) bool {
	return !IsEquals(a, b)
}
