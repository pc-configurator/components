package validation

func MinString(value *string, min int) bool {
	return len([]rune(*value)) >= min
}

type ErrorFields map[string]string

func (e ErrorFields) Error() string {
	return "validation error"
}
func (e ErrorFields) ConvertToMap() map[string]string {
	return e
}
