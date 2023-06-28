package TagUtils

import (
	"reflect"
	"strings"
)

// GetTagValue 获取给定字段的结构体标签值
func GetTagValue(fieldName string, structType reflect.Type, tagKey string) (string, bool) {
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if field.Name == fieldName {
			tag := field.Tag.Get(tagKey)
			return tag, true
		}
	}
	return "", false
}

// SetTagValue 设置给定字段的结构体标签值
func SetTagValue(fieldName string, structType reflect.Type, tagKey string, tagValue string) bool {
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		if field.Name != fieldName {
			continue
		}
		oldTag := field.Tag.Get(tagKey)
		if oldTag != "" {
			newTag := strings.ReplaceAll(field.Tag.Get(tagKey), oldTag, tagValue)
			// 只修改其中一个 tag
			tagString := strings.ReplaceAll(oldTag, oldTag, newTag)
			// 更新结构体标签值
			field.Tag = reflect.StructTag(tagString)
			return true
		}
		return false
	}
	return false
}
