package utils

// Map applies a function to all values
func Map(fn func(value interface{}) interface{}, values ...interface{}) []interface{} {
	var result []interface{}
	for _, value := range values {
		result = append(result, fn(value))
	}
	return result
}

// Truthy return true if the value is truthy
func Truthy(value interface{}) bool {
	switch typed := value.(type) {
	case int:
		return typed > 0
	case string:
		return len(typed) > 0
	case bool:
		return typed
	default:
		return false
	}
}

// First returns the first truthy value
func First(values ...interface{}) interface{} {
	for _, value := range values {
		if Truthy(value) {
			return value
		}
	}
	return nil
}

// FirstMap returns the first value whose function call result is truthy
func FirstMap(fn func(value interface{}) interface{}, values ...interface{}) interface{} {
	for _, value := range values {
		if Truthy(fn(value)) {
			return value
		}
	}
	return nil
}

// Any is true if any of the values is truthy
func Any(values ...interface{}) bool {
	return First(values...) != nil
}
