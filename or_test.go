package arm

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrDefault(t *testing.T) {
	require.Equal(t, "", Or[string]())
	require.Equal(t, 0, Or[int]())
}

func TestOr(t *testing.T) {
	require.Equal(t, 2, Or[int](0, 2))
	require.Equal(t, 2, Or[int](2, 1))
	require.Equal(t, "foo", Or[string]("", "", "foo"))
}
