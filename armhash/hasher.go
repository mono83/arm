package armhash

import "io"

// Hasher defines functions that can be used to produce hash
// This functions works with [io.Reader] for maximum versatility.
type Hasher[T any] func(io.Reader) (T, error)
