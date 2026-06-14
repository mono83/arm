package anies

import (
	"errors"
	"testing"

	"github.com/mono83/arm"
	"github.com/stretchr/testify/require"
)

type myBool bool

func TestToBool(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		assertBool(t, ToBool, true, true)
		assertBool(t, ToBool, false, false)
		assertBool(t, ToBool, arm.Ref(true), true)
		assertBool(t, ToBool, arm.Ref(false), false)
	})

	t.Run("signed integers", func(t *testing.T) {
		assertBool(t, ToBool, int(0), false)
		assertBool(t, ToBool, int(1), true)
		assertBool(t, ToBool, int8(0), false)
		assertBool(t, ToBool, int8(2), true)
		assertBool(t, ToBool, int16(0), false)
		assertBool(t, ToBool, int16(3), true)
		assertBool(t, ToBool, int32(0), false)
		assertBool(t, ToBool, int32(4), true)
		assertBool(t, ToBool, int64(0), false)
		assertBool(t, ToBool, int64(-5), true)
	})

	t.Run("unsigned integers", func(t *testing.T) {
		assertBool(t, ToBool, uint(0), false)
		assertBool(t, ToBool, uint(1), true)
		assertBool(t, ToBool, uint8(0), false)
		assertBool(t, ToBool, uint8(2), true)
		assertBool(t, ToBool, uint16(0), false)
		assertBool(t, ToBool, uint16(3), true)
		assertBool(t, ToBool, uint32(0), false)
		assertBool(t, ToBool, uint32(4), true)
		assertBool(t, ToBool, uint64(0), false)
		assertBool(t, ToBool, uint64(5), true)
	})

	t.Run("strings", func(t *testing.T) {
		for _, s := range []string{"true", "TRUE", " Yes ", "on", "1"} {
			assertBool(t, ToBool, s, true)
		}
		for _, s := range []string{"false", "no", "off", "0", "", "banana"} {
			assertBool(t, ToBool, s, false)
		}
	})

	t.Run("pointers", func(t *testing.T) {
		assertBool(t, ToBool, arm.Ref(1), true)
		assertBool(t, ToBool, arm.Ref(0), false)
		assertBool(t, ToBool, arm.Ref("yes"), true)
	})

	t.Run("nil", func(t *testing.T) {
		assertErr(t, ToBool, nil, ErrNilAny)
		var bp *bool
		assertErr(t, ToBool, bp, ErrNilAny)
		var ip *int
		assertErr(t, ToBool, ip, ErrNilAny)
	})

	t.Run("named types", func(t *testing.T) {
		assertBool(t, ToBool, myBool(true), true)
		assertBool(t, ToBool, myBool(false), false)
		assertBool(t, ToBool, myInt(0), false)
		assertBool(t, ToBool, myInt(3), true)
		assertBool(t, ToBool, myUint(0), false)
		assertBool(t, ToBool, myString("yes"), true)
		assertBool(t, ToBool, arm.Ref(myBool(true)), true)
	})

	t.Run("unsupported", func(t *testing.T) {
		assertErr(t, ToBool, 1.5, ErrUnsupported)
		assertErr(t, ToBool, struct{}{}, ErrUnsupported)
		assertErr(t, ToBool, myFloat(1), ErrUnsupported)
	})
}

func TestToBoolStrict(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		assertBool(t, ToBoolStrict, true, true)
		assertBool(t, ToBoolStrict, false, false)
		assertBool(t, ToBoolStrict, arm.Ref(true), true)
	})

	t.Run("nil", func(t *testing.T) {
		assertErr(t, ToBoolStrict, nil, ErrNilAny)
		var bp *bool
		assertErr(t, ToBoolStrict, bp, ErrNilAny)
	})

	t.Run("named types", func(t *testing.T) {
		assertBool(t, ToBoolStrict, myBool(true), true)
		assertBool(t, ToBoolStrict, arm.Ref(myBool(false)), false)
		assertErr(t, ToBoolStrict, myInt(1), ErrUnsupported)
		assertErr(t, ToBoolStrict, myString("true"), ErrUnsupported)
	})

	t.Run("unsupported", func(t *testing.T) {
		assertErr(t, ToBoolStrict, 1, ErrUnsupported)
		assertErr(t, ToBoolStrict, "true", ErrUnsupported)
	})
}

func assertBool(t *testing.T, fn func(any) (bool, error), in any, want bool) {
	t.Helper()
	got, err := fn(in)
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func assertErr(t *testing.T, fn func(any) (bool, error), in any, want error) {
	t.Helper()
	_, err := fn(in)
	require.True(t, errors.Is(err, want), "expected %v, got %v", want, err)
}
