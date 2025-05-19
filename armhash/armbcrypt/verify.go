package armbcrypt

import "golang.org/x/crypto/bcrypt"

// Verify performs verification of given hash
// against password and returns error on failure.
func Verify(hash string, password []byte) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), password)
}

// IsValid performs verification of given hash
// against password and returns error on failure.
func IsValid(hash string, password []byte) bool {
	return Verify(hash, password) == nil
}
