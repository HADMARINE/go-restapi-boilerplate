package util

// IsNil checks nil of any type
func IsNil(i interface{}) bool {
	return i == nil
}

// IsNilArray verifies nil of array, any of it has nil, verify it with true value
func IsNilArray(values ...interface{}) bool {
	for _, val := range values {
		if IsNil(val) == true {
			return true
		}
	}
	return false
}
