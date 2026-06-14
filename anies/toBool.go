package anies

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/mono83/arm"
)

// ToBool performs a lenient conversion of an arbitrary value into a bool.
//
// Supported inputs:
//   - bool is returned as-is.
//   - Signed and unsigned integers convert to false when zero, true otherwise.
//   - Strings are trimmed and matched case-insensitively against the truthy
//     tokens "TRUE", "YES", "ON" and "1"; any other string yields false.
//   - Pointers are dereferenced and their target converted recursively.
//   - Named scalar types (e.g. `type Flag bool`) are converted by their base.
//
// A nil input (including a typed nil) returns ErrNilAny. Any other type
// returns ErrUnsupported.
func ToBool(a any) (bool, error) {
	switch x := a.(type) {
	case bool:
		return x, nil
	case int:
		return x != 0, nil
	case int8:
		return x != 0, nil
	case int16:
		return x != 0, nil
	case int32:
		return x != 0, nil
	case int64:
		return x != 0, nil
	case uint:
		return x != 0, nil
	case uint8:
		return x != 0, nil
	case uint16:
		return x != 0, nil
	case uint32:
		return x != 0, nil
	case uint64:
		return x != 0, nil
	case string:
		s := strings.ToUpper(strings.TrimSpace(x))
		return s == "TRUE" || s == "YES" || s == "ON" || s == "1", nil
	case nil:
		return false, ErrNilAny
	default:
		return derefBool(a, ToBool)
	}
}

// ToBoolStrict converts a value into a bool accepting only boolean inputs.
//
// Only bool (and pointers to or named types over bool) are accepted. A nil
// input (including a typed nil) returns ErrNilAny; every other type returns
// ErrUnsupported.
func ToBoolStrict(a any) (bool, error) {
	switch x := a.(type) {
	case bool:
		return x, nil
	case nil:
		return false, ErrNilAny
	default:
		return derefBool(a, ToBoolStrict)
	}
}

// derefBool handles the cold path of the bool converters: typed nils yield
// ErrNilAny, pointers and named scalars are unwrapped and re-run through conv,
// and everything else is unsupported. Recursing through conv keeps each
// converter's own semantics: a named integer resolves under ToBool but is
// rejected by ToBoolStrict. Keeping this out of the type switch lets concrete
// inputs resolve without touching reflection.
func derefBool(a any, conv func(any) (bool, error)) (bool, error) {
	if arm.IsNil(a) {
		return false, ErrNilAny
	}
	rv := reflect.ValueOf(a)
	if rv.Kind() == reflect.Ptr {
		return conv(rv.Elem().Interface())
	}
	if base, ok := basic(rv); ok {
		return conv(base)
	}
	return false, fmt.Errorf("%T %w", a, ErrUnsupported)
}
