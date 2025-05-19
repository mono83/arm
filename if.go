package arm

// If is simple generic ternary operator implementation
func If[T any](condition bool, a T, b T) T {
	if condition {
		return a
	}
	return b
}
