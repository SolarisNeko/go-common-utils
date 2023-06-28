package stream

import (
	"github.com/SolarisNeko/go-common-utils/generic"
)

type T generic.Type
type Number generic.Number

// Stream 泛型切片
type Stream []T

// Filter 方法用于过滤满足条件的元素
func (stream Stream) Filter(predicate func(T) bool) Stream {
	result := make(Stream, 0)
	for _, item := range stream {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map 方法用于对每个元素执行映射操作
func (stream Stream) Map(mapper func(T) T) Stream {
	result := make(Stream, len(stream))
	for i, item := range stream {
		result[i] = mapper(item)
	}
	return result
}

// Reduce 方法用于将元素缩减为单个结果
func (stream Stream) Reduce(initial T, reducer func(T, T) T) T {
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
