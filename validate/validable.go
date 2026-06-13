package validate

// Validable is implemented by types that can validate themselves.
type Validable interface {
	// Validate reports whether the value is valid, returning a non-nil error
	// describing the failure otherwise.
	Validate() error
}
