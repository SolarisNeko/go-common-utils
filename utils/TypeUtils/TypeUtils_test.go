package TypeUtils

import (
	"testing"

	"github.com/SolarisNeko/go-common-utils/utils/ArrayUtils"
)

type User struct {
	userId   int
	username string
}

func TestGetFieldNames(t *testing.T) {
	type args struct {
		obj interface{}
	}

	fieldNames := GetFieldNames(User{})

	expected := []string{"userId", "username"}

	for _, fieldName := range fieldNames {
		if ArrayUtils.Contains(expected, fieldName) {
		} else {
			t.Errorf("Add(2, 3) returned %s, expected %s", fieldNames, expected)
		}

	}
}

func TestGetMethodNames(t *testing.T) {
}
