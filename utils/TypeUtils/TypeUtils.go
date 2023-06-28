package TypeUtils

import (
	"reflect"
)

// GetFieldNames 获取结构体或类的字段名列表
func GetFieldNames(obj interface{}) []string {
	objType := reflect.TypeOf(obj)
	if objType.Kind() != reflect.Struct {
		return nil
	}

	fieldNameArray := make([]string, 0)
	numFields := objType.NumField()
	for i := 0; i < numFields; i++ {
		field := objType.Field(i)
		fieldNameArray = append(fieldNameArray, field.Name)
	}

	return fieldNameArray
}

// GetMethodNames 获取结构体或类的方法名列表
func GetMethodNames(obj interface{}) []string {
	objType := reflect.TypeOf(obj)
	if objType.Kind() != reflect.Struct {
		return nil
	}

	methodNameArray := make([]string, 0)
	numMethods := objType.NumMethod()
	for i := 0; i < numMethods; i++ {
		method := objType.Method(i)
		methodNameArray = append(methodNameArray, method.Name)
	}

	return methodNameArray
}
