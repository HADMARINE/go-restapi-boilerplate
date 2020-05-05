package util

import "reflect"

// IsNil checks nil of any type
func IsNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}

// IsNilArray verifies nil of array, any of it has nil
func IsNilArray(want bool, values ...interface{}) bool {
	var flag bool = false
	for val := range values {
		if IsNil(val) == !want {
			flag = true
		}
	}
	return !flag
}
