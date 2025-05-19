package arm

// Must takes tuple of value with error and returns only value
// if error is nil. If error not nil - panics it.
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}
