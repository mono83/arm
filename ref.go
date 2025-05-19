package arm

// Ref returns reference to given value
func Ref[T any](t T) *T { return &t }
