package armbcrypt

import (
	"bytes"
	"github.com/mono83/arm/armhash"
	"golang.org/x/crypto/bcrypt"
	"io"
)

// NewHasher constructs new hasher with given cost.
// If zero or negative cost given, returns hasher
// with default bcrypt cost (10 for the moment).
func NewHasher(cost int) armhash.Hasher[string] {
	if cost < 1 {
		cost = bcrypt.DefaultCost
	}
	return func(reader io.Reader) (string, error) {
		buf := new(bytes.Buffer)
		_, err := io.Copy(buf, reader)
		if err != nil {
			return "", err
		}

		hash, err := bcrypt.GenerateFromPassword(buf.Bytes(), cost)
		if err != nil {
			return "", err
		}

		return string(hash), nil
	}
}

// NewDefaultHasher construct new hasher with default
// bcrypt cost (10 for the moment).
func NewDefaultHasher() armhash.Hasher[string] { return NewHasher(-1) }
