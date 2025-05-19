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
