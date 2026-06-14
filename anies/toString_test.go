package anies

import (
	"errors"
	"testing"

	"github.com/mono83/arm"
	"github.com/stretchr/testify/require"
)

type (
	myString  string
	myFloat32 float32
)

type stringer struct{}

func (stringer) String() string { return "STRINGER" }

type ptrStringer struct{}

func (*ptrStringer) String() string { return "PTR" }

func TestToString(t *testing.T) {
	t.Run("strings and bytes", func(t *testing.T) {
		assertStr(t, "hello", "hello")
		assertStr(t, []byte("bytes"), "bytes")
		assertStr(t, "", "")
	})

	t.Run("bool", func(t *testing.T) {
		assertStr(t, true, "true")
		assertStr(t, false, "false")
	})

	t.Run("signed integers", func(t *testing.T) {
		assertStr(t, int(-7), "-7")
		assertStr(t, int8(8), "8")
		assertStr(t, int16(16), "16")
		assertStr(t, int32(32), "32")
		assertStr(t, int64(64), "64")
	})

	t.Run("unsigned integers", func(t *testing.T) {
		assertStr(t, uint(1), "1")
		assertStr(t, uint8(8), "8")
		assertStr(t, uint16(16), "16")
		assertStr(t, uint32(32), "32")
		assertStr(t, uint64(64), "64")
	})

	t.Run("floats", func(t *testing.T) {
		assertStr(t, float32(1.5), "1.5")
		assertStr(t, float64(-2.25), "-2.25")
	})

	t.Run("error and stringer", func(t *testing.T) {
		assertStr(t, errors.New("boom"), "boom")
		assertStr(t, stringer{}, "STRINGER")
		assertStr(t, &ptrStringer{}, "PTR")
	})

	t.Run("pointers", func(t *testing.T) {
		assertStr(t, arm.Ref(99), "99")
		assertStr(t, arm.Ref("x"), "x")
		assertStr(t, arm.Ref(true), "true")
	})

	t.Run("nil", func(t *testing.T) {
		_, err := ToString(nil)
		require.ErrorIs(t, err, ErrNilAny)

		var ip *int
		_, err = ToString(ip)
		require.ErrorIs(t, err, ErrNilAny)

		var ps *ptrStringer
		_, err = ToString(ps)
		require.ErrorIs(t, err, ErrNilAny)
	})

	t.Run("named types", func(t *testing.T) {
		assertStr(t, myString("hi"), "hi")
		assertStr(t, myInt(-7), "-7")
		assertStr(t, myUint(8), "8")
		assertStr(t, myBool(true), "true")
		assertStr(t, myFloat(-2.25), "-2.25")
		assertStr(t, myFloat32(1.5), "1.5")
		assertStr(t, arm.Ref(myString("x")), "x")
	})

	t.Run("unsupported", func(t *testing.T) {
		_, err := ToString(struct{}{})
		require.ErrorIs(t, err, ErrUnsupported)
	})
}

func assertStr(t *testing.T, in any, want string) {
	t.Helper()
	got, err := ToString(in)
	require.NoError(t, err)
	require.Equal(t, want, got)
}
