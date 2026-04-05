package arm

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var errProvide = errors.New("provider error")

func ok[T any](v T) func() (T, error) { return func() (T, error) { return v, nil } }
func bad[T any]() func() (T, error)   { return func() (T, error) { var z T; return z, errProvide } }

func TestAllOfProvided2(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		v1, v2, err := AllOfProvided2(ok(1), ok("a"))
		require.NoError(t, err)
		require.Equal(t, 1, v1)
		require.Equal(t, "a", v2)
	})
	t.Run("f1 error", func(t *testing.T) {
		_, _, err := AllOfProvided2(bad[int](), ok("a"))
		require.ErrorIs(t, err, errProvide)
	})
	t.Run("f2 error", func(t *testing.T) {
		_, _, err := AllOfProvided2(ok(1), bad[string]())
		require.ErrorIs(t, err, errProvide)
	})
	t.Run("f1 nil", func(t *testing.T) {
		_, _, err := AllOfProvided2[int, string](nil, ok("a"))
		require.EqualError(t, err, "f1 is nil")
	})
	t.Run("f2 nil", func(t *testing.T) {
		_, _, err := AllOfProvided2[int, string](ok(1), nil)
		require.EqualError(t, err, "f2 is nil")
	})
}

func TestAllOfProvided3(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		v1, v2, v3, err := AllOfProvided3(ok(1), ok("a"), ok(true))
		require.NoError(t, err)
		require.Equal(t, 1, v1)
		require.Equal(t, "a", v2)
		require.True(t, v3)
	})
	t.Run("f2 error", func(t *testing.T) {
		_, _, _, err := AllOfProvided3(ok(1), bad[string](), ok(true))
		require.ErrorIs(t, err, errProvide)
	})
	t.Run("f3 nil", func(t *testing.T) {
		_, _, _, err := AllOfProvided3[int, string, bool](ok(1), ok("a"), nil)
		require.EqualError(t, err, "f3 is nil")
	})
}

func TestAllOfProvided4(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		v1, v2, v3, v4, err := AllOfProvided4(ok(1), ok("a"), ok(true), ok(3.14))
		require.NoError(t, err)
		require.Equal(t, 1, v1)
		require.Equal(t, "a", v2)
		require.True(t, v3)
		require.InDelta(t, 3.14, v4, 1e-9)
	})
	t.Run("f3 error", func(t *testing.T) {
		_, _, _, _, err := AllOfProvided4(ok(1), ok("a"), bad[bool](), ok(3.14))
		require.ErrorIs(t, err, errProvide)
	})
	t.Run("f4 nil", func(t *testing.T) {
		_, _, _, _, err := AllOfProvided4[int, string, bool, float64](ok(1), ok("a"), ok(true), nil)
		require.EqualError(t, err, "f4 is nil")
	})
}

func TestAllOfProvided5(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		v1, v2, v3, v4, v5, err := AllOfProvided5(ok(1), ok("a"), ok(true), ok(3.14), ok(uint(7)))
		require.NoError(t, err)
		require.Equal(t, 1, v1)
		require.Equal(t, "a", v2)
		require.True(t, v3)
		require.InDelta(t, 3.14, v4, 1e-9)
		require.Equal(t, uint(7), v5)
	})
	t.Run("f4 error", func(t *testing.T) {
		_, _, _, _, _, err := AllOfProvided5(ok(1), ok("a"), ok(true), bad[float64](), ok(uint(7)))
		require.ErrorIs(t, err, errProvide)
	})
	t.Run("f5 nil", func(t *testing.T) {
		_, _, _, _, _, err := AllOfProvided5[int, string, bool, float64, uint](ok(1), ok("a"), ok(true), ok(3.14), nil)
		require.EqualError(t, err, "f5 is nil")
	})
}
