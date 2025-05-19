package armhash

import (
	"crypto/sha256"
	"io"
)

// SHA256 is a plain SHA-256 implementation
func SHA256(r io.Reader) ([]byte, error) {
	h := sha256.New()
	if _, err := io.Copy(h, r); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}
