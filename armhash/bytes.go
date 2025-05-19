package armhash

import "bytes"

// Bytes hashes given bytes using provided hasher.
func Bytes[T any](hash Hasher[T], bts []byte) (T, error) {
	return hash(bytes.NewReader(bts))
}
