package validation

import "strconv"

func MinString(value *string, min int) bool {
	return len([]rune(*value)) >= min
}

func IsNumericString(value string) bool {
	if _, err := strconv.ParseInt(value, 10, 64); err != nil {
		return false
	}

	return true
}

type ErrorFields map[string]string

func (e ErrorFields) Error() string {
	return "validation error"
}
func (e ErrorFields) ConvertToMap() map[string]string {
	return e
}
