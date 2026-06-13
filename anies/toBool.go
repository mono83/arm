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
// Only bool (and pointers to bool, dereferenced) are accepted. A nil input
// (including a typed nil) returns ErrNilAny; every other type returns
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
// ErrNilAny, pointers are dereferenced and re-run through conv, and everything
// else is unsupported. Keeping this out of the type switch lets concrete
// inputs resolve without touching reflection.
func derefBool(a any, conv func(any) (bool, error)) (bool, error) {
	if arm.IsNil(a) {
		return false, ErrNilAny
	}
	if rv := reflect.ValueOf(a); rv.Kind() == reflect.Ptr {
		return conv(rv.Elem().Interface())
	}
	return false, fmt.Errorf("%T %w", a, ErrUnsupported)
}
