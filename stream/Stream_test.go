package stream

import (
	"fmt"
	"testing"
	_ "testing"

	"github.com/SolarisNeko/go-common-utils/generic"
)

func TestStream_Filter(t *testing.T) {
	stream := Stream{1, 2, 3, 4, 5}

	// 使用 Filter 过滤偶数
	predicate := func(num interface{}) bool {
		if number, ok := num.(int); ok {
			return number%2 == 0
		}
		return false
	}
	// 使用 Map 将每个元素加倍
	mapper := func(n interface{}) interface{} {
		toNumber := generic.ToNumber(n)
		number := generic.ToNumber(2)
		return toNumber * number
	}

	reducer := func(n1 interface{}, n2 interface{}) interface{} {
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

type Person struct {
	Name   string
	Age    int
	Gender string
}

func TestStream_GroupBy(t *testing.T) {
	people := Stream{
		Person{Name: "Alice", Age: 30, Gender: "Female"},
		Person{Name: "Bob", Age: 28, Gender: "Male"},
		Person{Name: "Charlie", Age: 35, Gender: "Male"},
		Person{Name: "Eve", Age: 30, Gender: "Female"},
		Person{Name: "David", Age: 28, Gender: "Male"},
	}

	// 分组并聚合为 map[string]interface{}[]
	var grouped = people.GroupBy(
		func(p interface{}) string {
			return p.(Person).Gender
		},
	)

	for key, values := range grouped {
		fmt.Printf("基于性别分组, 性别 = %s\n", key)
		for _, value := range values {
			person := value.(Person)
			fmt.Printf("姓名：%s，年龄：%d\n", person.Name, person.Age)
		}
		fmt.Println()
	}

}
