package arm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsNil(t *testing.T) {
	t.Run("untyped nil", func(t *testing.T) {
		require.True(t, IsNil(nil))
	})

	t.Run("typed nil pointers", func(t *testing.T) {
		var bp *bool
		var ip *int
		var sp *string
		require.True(t, IsNil(bp))
		require.True(t, IsNil(ip))
		require.True(t, IsNil(sp))
	})

	t.Run("nil reference kinds", func(t *testing.T) {
		var m map[string]int
		var s []int
		var ch chan int
		var fn func()
		require.True(t, IsNil(m))
		require.True(t, IsNil(s))
		require.True(t, IsNil(ch))
		require.True(t, IsNil(fn))
	})

	t.Run("nil interface value", func(t *testing.T) {
		var err error
		require.True(t, IsNil(err))
	})

	t.Run("non-nil pointer", func(t *testing.T) {
		b := true
		require.False(t, IsNil(&b))
	})

	t.Run("non-nilable values", func(t *testing.T) {
		require.False(t, IsNil(0))
		require.False(t, IsNil(""))
		require.False(t, IsNil(false))
		require.False(t, IsNil(struct{}{}))
		require.False(t, IsNil([0]int{}))
	})

	t.Run("non-nil reference kinds", func(t *testing.T) {
		require.False(t, IsNil(map[string]int{}))
		require.False(t, IsNil([]int{}))
		require.False(t, IsNil(make(chan int)))
		require.False(t, IsNil(func() {}))
	})
}
