package anies

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/mono83/arm"
)

// ToString performs a lenient conversion of an arbitrary value into a string.
//
// Supported inputs:
//   - string is returned as-is; []byte is converted directly.
//   - bool becomes "true" or "false".
//   - Signed and unsigned integers are formatted in base 10.
//   - Floats use the shortest representation that round-trips.
//   - Values implementing error or fmt.Stringer use Error()/String().
//   - Pointers are dereferenced and their target converted recursively.
//
// A nil input (including a typed nil) returns ErrNilAny. Any other type
// returns ErrUnsupported.
func ToString(a any) (string, error) {
	switch x := a.(type) {
	case string:
		return x, nil
	case []byte:
		return string(x), nil
	case bool:
		return strconv.FormatBool(x), nil
	case int:
		return strconv.FormatInt(int64(x), 10), nil
	case int8:
		return strconv.FormatInt(int64(x), 10), nil
	case int16:
		return strconv.FormatInt(int64(x), 10), nil
	case int32:
		return strconv.FormatInt(int64(x), 10), nil
	case int64:
		return strconv.FormatInt(x, 10), nil
	case uint:
		return strconv.FormatUint(uint64(x), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(x), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(x), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(x), 10), nil
	case uint64:
		return strconv.FormatUint(x, 10), nil
	case float32:
		return strconv.FormatFloat(float64(x), 'g', -1, 32), nil
	case float64:
		return strconv.FormatFloat(x, 'g', -1, 64), nil
	case error:
		// A typed nil stored in the interface would panic on Error().
		if arm.IsNil(x) {
			return "", ErrNilAny
		}
		return x.Error(), nil
	case fmt.Stringer:
		if arm.IsNil(x) {
			return "", ErrNilAny
		}
		return x.String(), nil
	case nil:
		return "", ErrNilAny
	default:
		if arm.IsNil(a) {
			return "", ErrNilAny
		}
		if rv := reflect.ValueOf(a); rv.Kind() == reflect.Ptr {
			return ToString(rv.Elem().Interface())
		}
		return "", fmt.Errorf("%T %w", a, ErrUnsupported)
	}
}
