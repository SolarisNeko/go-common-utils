# go sugar

# Target
1. common 工具包
2. Java Stream / C# LinQ 风格带到 Go
3. 追加一个泛型特性. 弥补 Go 还没出泛型时的需要

# Feature 特性

1. Java Stream -> Go Stream
2. Java 泛型 <T> -> Go T

# Download
```shell
go get -u github.com/SolarisNeko/go-common-utils

```

# Use
## Stream
```go
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

```