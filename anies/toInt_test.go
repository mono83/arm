package anies

import (
	"errors"
	"math"
	"testing"

	"github.com/mono83/arm"
	"github.com/stretchr/testify/require"
)

func assertErrInt(t *testing.T, fn func(any) (int, error), in any, want error) {
	t.Helper()
	_, err := fn(in)
	require.True(t, errors.Is(err, want), "expected %v, got %v", want, err)
}

func TestToInt(t *testing.T) {
	t.Run("signed integers", func(t *testing.T) {
		assertInt(t, ToInt, int(0), 0)
		assertInt(t, ToInt, int(-7), -7)
		assertInt(t, ToInt, int8(8), 8)
		assertInt(t, ToInt, int16(16), 16)
		assertInt(t, ToInt, int32(32), 32)
		assertInt(t, ToInt, int64(64), 64)
	})

	t.Run("unsigned integers", func(t *testing.T) {
		assertInt(t, ToInt, uint(1), 1)
		assertInt(t, ToInt, uint8(8), 8)
		assertInt(t, ToInt, uint16(16), 16)
		assertInt(t, ToInt, uint32(32), 32)
		assertInt(t, ToInt, uint64(64), 64)
	})

	t.Run("floats truncate", func(t *testing.T) {
		assertInt(t, ToInt, float32(3.9), 3)
		assertInt(t, ToInt, float64(-2.9), -2)
	})

	t.Run("bool", func(t *testing.T) {
		assertInt(t, ToInt, true, 1)
		assertInt(t, ToInt, false, 0)
	})

	t.Run("strings", func(t *testing.T) {
		assertInt(t, ToInt, "42", 42)
		assertInt(t, ToInt, " -5 ", -5)
		assertErrInt(t, ToInt, "banana", ErrUnsupported)
		assertErrInt(t, ToInt, "3.14", ErrUnsupported)
	})

	t.Run("pointers", func(t *testing.T) {
		assertInt(t, ToInt, arm.Ref(99), 99)
		assertInt(t, ToInt, arm.Ref(int64(123)), 123)
		assertInt(t, ToInt, arm.Ref(true), 1)
		assertInt(t, ToInt, arm.Ref("7"), 7)
	})

	t.Run("nil", func(t *testing.T) {
		assertErrInt(t, ToInt, nil, ErrNilAny)
		var ip *int
		assertErrInt(t, ToInt, ip, ErrNilAny)
	})

	t.Run("unsupported", func(t *testing.T) {
		assertErrInt(t, ToInt, struct{}{}, ErrUnsupported)
	})

	t.Run("overflow", func(t *testing.T) {
		assertErrInt(t, ToInt, uint64(math.MaxUint64), ErrOverflow)
		assertErrInt(t, ToInt, math.MaxFloat64, ErrOverflow)
		assertErrInt(t, ToInt, math.NaN(), ErrOverflow)
		assertErrInt(t, ToInt, "99999999999999999999999", ErrOverflow)
		assertInt(t, ToInt, math.MaxInt, math.MaxInt)
	})
}

func TestToIntStrict(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		assertInt(t, ToIntStrict, int(5), 5)
		assertInt(t, ToIntStrict, int64(64), 64)
		assertInt(t, ToIntStrict, uint8(8), 8)
		assertInt(t, ToIntStrict, arm.Ref(7), 7)
	})

	t.Run("nil", func(t *testing.T) {
		assertErrInt(t, ToIntStrict, nil, ErrNilAny)
		var ip *int
		assertErrInt(t, ToIntStrict, ip, ErrNilAny)
	})

	t.Run("unsupported", func(t *testing.T) {
		assertErrInt(t, ToIntStrict, 1.5, ErrUnsupported)
		assertErrInt(t, ToIntStrict, true, ErrUnsupported)
		assertErrInt(t, ToIntStrict, "42", ErrUnsupported)
	})

	t.Run("overflow", func(t *testing.T) {
		assertErrInt(t, ToIntStrict, uint64(math.MaxUint64), ErrOverflow)
		assertErrInt(t, ToIntStrict, uint(math.MaxUint), ErrOverflow)
		assertInt(t, ToIntStrict, uint64(math.MaxInt), math.MaxInt)
	})
}

func assertInt(t *testing.T, fn func(any) (int, error), in any, want int) {
	t.Helper()
	got, err := fn(in)
	require.NoError(t, err)
	require.Equal(t, want, got)
}
