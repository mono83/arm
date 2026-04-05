package arm

// Or returns first not-default value among all given candidates.
// If candidates slice is empty it returns default value for T.
func Or[T comparable](candidates ...T) (out T) {
	for _, t := range candidates {
		if t != out {
			// Candidate value is not default one
			return t
		}
	}
	return
}

// OrUnref returns value of first argument if it's not nil
// otherwise returns second argument.
func OrUnref[T any](v *T, def T) T {
	if v == nil {
		return def
	}
	return *v
}

// OrProvide returns t if it is not the zero value for T,
// otherwise it calls provide and returns its result.
func OrProvide[T comparable](t T, provide func() T) (out T) {
	if t != out {
		return t
	}
	return provide()
}
