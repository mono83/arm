package arm

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMust(t *testing.T) {
	require.Equal(t, "foo", Must("foo", nil))
	require.Panics(t, func() {
		Must("foo", errors.New("foo"))
	})
}
