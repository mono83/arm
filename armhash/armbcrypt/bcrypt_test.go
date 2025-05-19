package armbcrypt

import (
	"github.com/mono83/arm/armhash"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBCrypt(t *testing.T) {
	h := NewHasher(8) // Using small cost for faster testing
	if hash, err := armhash.String(h, "foo"); assert.NoError(t, err) {
		assert.True(t, IsValid(hash, []byte("foo")), "Hash verification failed")
		assert.False(t, IsValid(hash, []byte("bar")), "Hash verification failed")
		assert.True(t, strings.HasPrefix(hash, "$2a$08$"))
	}
}
