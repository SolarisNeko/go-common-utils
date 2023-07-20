package stream

import (
	"github.com/SolarisNeko/go-common-utils/generic"
)

type Number generic.Number

// Stream 泛型切片
type Stream []interface{}

// GroupResult 结构表示分组后的结果
type GroupResult struct {
	Key    string
	Values []interface{}
}

// Filter 方法用于过滤满足条件的元素
func (stream Stream) Filter(predicate func(interface{}) bool) Stream {
	result := make(Stream, 0)
	for _, item := range stream {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map 方法用于对每个元素执行映射操作
func (stream Stream) Map(mapper func(interface{}) interface{}) Stream {
	result := make(Stream, len(stream))
	for i, item := range stream {
		result[i] = mapper(item)
	}
	return result
}

// Reduce 方法用于将元素缩减为单个结果
func (stream Stream) Reduce(initial interface{}, reducer func(interface{}, interface{}) interface{}) interface{} {
	result := initial
	for _, item := range stream {
		result = reducer(result, item)
	}
	return result
}

// FlatMap 对切片中的元素进行映射和展开操作
func (stream Stream) FlatMap(mapper func(interface{}) Stream) Stream {
	result := make(Stream, 0)
	for _, item := range stream {
		mapped := mapper(item)
		result = append(result, mapped...)
	}
	return result
}

// First 返回切片中的第一个元素，如果切片为空则返回 nil
func (stream Stream) First() interface{} {
	if len(stream) > 0 {
		return stream[0]
	}
	return nil
}

func Count(stream Stream) int {
	count := 0
	for range stream {
		count++
	}
	return count
}

// GroupBy 方法用于将元素按照条件进行分组，并返回 map[string]interface{}[]
func (stream Stream) GroupBy(keySelector func(interface{}) string) map[string][]interface{} {
	groups := make(map[string][]interface{})

	for _, item := range stream {
		key := keySelector(item)
		groups[key] = append(groups[key], item)
	}

	return groups
}
