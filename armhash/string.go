package armhash

import "strings"

// String hashes given bytes using provided hasher.
func String[T any](hash Hasher[T], s string) (T, error) {
	return hash(strings.NewReader(s))
}
