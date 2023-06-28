package stream

import (
	"testing"

	"github.com/SolarisNeko/go-common-utils/generic"
)

func TestStream_Filter(t *testing.T) {
	stream := Stream{1, 2, 3, 4, 5}

	// 使用 Filter 过滤偶数
	predicate := func(num T) bool {
		if number, ok := num.(int); ok {
			return number%2 == 0
		}
		return false
	}
	// 使用 Map 将每个元素加倍
	mapper := func(n T) T {
		toNumber := generic.ToNumber(n)
		number := generic.ToNumber(2)
		return toNumber * number
	}

	reducer := func(n1 T, n2 T) T {
		return generic.ToNumber(n1) + generic.ToNumber(n2)
	}

	// 输出: 12
	sum := stream.
		Filter(predicate).
		Map(mapper).
		Reduce(0, reducer)

	expect := Number(12)
	if expect == sum {
		t.Errorf("value not equals 12, value = %v", sum)
	}
}
