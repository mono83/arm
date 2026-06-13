package validate

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

var errInvalid = errors.New("invalid")

// failing is a Validable that always reports an error.
type failing struct{}

func (failing) Validate() error { return errInvalid }

// passing is a Validable that always succeeds.
type passing struct{}

func (passing) Validate() error { return nil }

func TestReflectiveValidatorNil(t *testing.T) {
	t.Run("disallowed", func(t *testing.T) {
		require.ErrorIs(t, ReflectiveValidator{}.Validate(nil), ErrNil)
	})
	t.Run("allowed", func(t *testing.T) {
		require.NoError(t, ReflectiveValidator{AllowNil: true}.Validate(nil))
	})
	t.Run("typed nil pointer", func(t *testing.T) {
		var p *passing
		require.ErrorIs(t, ReflectiveValidator{}.Validate(p), ErrNil)
	})
}

func TestReflectiveValidatorValidable(t *testing.T) {
	require.ErrorIs(t, ReflectiveValidator{}.Validate(failing{}), errInvalid)
	require.NoError(t, ReflectiveValidator{}.Validate(passing{}))
}

func TestReflectiveValidatorPointer(t *testing.T) {
	require.ErrorIs(t, ReflectiveValidator{}.Validate(&failing{}), errInvalid)
}

func TestReflectiveValidatorScalar(t *testing.T) {
	require.NoError(t, ReflectiveValidator{}.Validate(42))
	require.NoError(t, ReflectiveValidator{}.Validate("text"))
}

func TestReflectiveValidatorSlice(t *testing.T) {
	t.Run("recurses", func(t *testing.T) {
		err := ReflectiveValidator{}.Validate([]Validable{passing{}, failing{}})
		require.ErrorIs(t, err, errInvalid)
		require.Contains(t, err.Error(), "validate: [1]: invalid")
	})
	t.Run("skipped", func(t *testing.T) {
		err := ReflectiveValidator{SkipSliceValues: true}.Validate([]Validable{failing{}})
		require.NoError(t, err)
	})
}

func TestReflectiveValidatorMap(t *testing.T) {
	m := map[string]Validable{"a": failing{}}

	t.Run("values", func(t *testing.T) {
		require.ErrorIs(t, ReflectiveValidator{}.Validate(m), errInvalid)
	})
	t.Run("values skipped", func(t *testing.T) {
		require.NoError(t, ReflectiveValidator{SkipMapValues: true}.Validate(m))
	})
	t.Run("failing key", func(t *testing.T) {
		mk := map[failing]string{{}: "x"}
		require.ErrorIs(t, ReflectiveValidator{}.Validate(mk), errInvalid)
		require.NoError(t, ReflectiveValidator{SkipMapKeys: true}.Validate(mk))
	})
}

func TestReflectiveValidatorStruct(t *testing.T) {
	type box struct {
		Public  Validable
		private Validable
	}

	t.Run("recurses into exported fields", func(t *testing.T) {
		err := ReflectiveValidator{}.Validate(box{Public: failing{}})
		require.ErrorIs(t, err, errInvalid)
		require.Contains(t, err.Error(), "validate: Public: invalid")
	})
	t.Run("reports nested location", func(t *testing.T) {
		type list struct {
			Items []Validable
		}
		err := ReflectiveValidator{}.Validate(list{Items: []Validable{passing{}, failing{}}})
		require.ErrorIs(t, err, errInvalid)
		require.Contains(t, err.Error(), "validate: Items: validate: [1]: invalid")
	})
	t.Run("skips unexported fields", func(t *testing.T) {
		require.NoError(t, ReflectiveValidator{}.Validate(box{Public: passing{}, private: failing{}}))
	})
	t.Run("skipped", func(t *testing.T) {
		err := ReflectiveValidator{SkipStruct: true}.Validate(box{Public: failing{}})
		require.NoError(t, err)
	})
}

func TestReflectiveValidatorArray(t *testing.T) {
	require.ErrorIs(t, ReflectiveValidator{}.Validate([2]Validable{passing{}, failing{}}), errInvalid)
}

func TestReflectiveValidatorNilElement(t *testing.T) {
	t.Run("disallowed", func(t *testing.T) {
		require.ErrorIs(t, ReflectiveValidator{}.Validate([]Validable{nil}), ErrNil)
	})
	t.Run("allowed", func(t *testing.T) {
		require.NoError(t, ReflectiveValidator{AllowNil: true}.Validate([]Validable{nil}))
	})
}

func TestReflectiveValidatorCycle(t *testing.T) {
	t.Run("pointer", func(t *testing.T) {
		type node struct {
			Self *node
			Leaf Validable
		}
		n := &node{Leaf: failing{}}
		n.Self = n
		require.ErrorIs(t, ReflectiveValidator{}.Validate(n), errInvalid)
	})
	t.Run("slice", func(t *testing.T) {
		s := make([]any, 1)
		s[0] = s
		require.NoError(t, ReflectiveValidator{}.Validate(s))
	})
	t.Run("map", func(t *testing.T) {
		m := map[string]any{}
		m["self"] = m
		require.NoError(t, ReflectiveValidator{}.Validate(m))
	})
}

func TestReflectiveValidatorSharedReference(t *testing.T) {
	type box struct {
		A, B *failing
	}
	shared := &failing{}
	err := ReflectiveValidator{}.Validate(box{A: shared, B: shared})
	require.ErrorIs(t, err, errInvalid)
	// The shared pointer is on a sibling branch, not the path, so both fields
	// are validated rather than one being mistaken for a cycle.
	require.Contains(t, err.Error(), "validate: A: invalid")
	require.Contains(t, err.Error(), "validate: B: invalid")
}
