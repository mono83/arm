package anies

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/mono83/arm"
)

// ToInt performs a lenient conversion of an arbitrary value into an int.
//
// Supported inputs:
//   - Signed and unsigned integers are converted directly.
//   - Floats are truncated towards zero.
//   - bool converts to 1 (true) or 0 (false).
//   - Strings are trimmed and parsed as a base-10 integer; an unparsable
//     string returns ErrUnsupported.
//   - Pointers are dereferenced and their target converted recursively.
//
// A value that does not fit into an int (a large unsigned integer, an
// out-of-range float or numeric string, or - on 32-bit platforms - a large
// int64) returns ErrOverflow. A nil input (including a typed nil) returns
// ErrNilAny. Any other type returns ErrUnsupported.
func ToInt(a any) (int, error) {
	switch x := a.(type) {
	case float32:
		return floatToInt(float64(x))
	case float64:
		return floatToInt(x)
	case bool:
		if x {
			return 1, nil
		}
		return 0, nil
	case string:
		n, err := strconv.ParseInt(strings.TrimSpace(x), 10, 0)
		if err != nil {
			if errors.Is(err, strconv.ErrRange) {
				return 0, fmt.Errorf("%q: %w", x, ErrOverflow)
			}
			return 0, fmt.Errorf("%q: %w", x, ErrUnsupported)
		}
		return int(n), nil
	case nil:
		return 0, ErrNilAny
	default:
		// Pointers recurse leniently (so e.g. *bool is supported); every
		// remaining type - all integers included - is handled by ToIntStrict.
		if arm.IsNil(a) {
			return 0, ErrNilAny
		}
		if rv := reflect.ValueOf(a); rv.Kind() == reflect.Ptr {
			return ToInt(rv.Elem().Interface())
		}
		return ToIntStrict(a)
	}
}

// ToIntStrict converts a value into an int accepting only integer inputs.
//
// Only signed and unsigned integer types (and pointers to them) are accepted;
// floats, bools and strings are rejected. A value that does not fit into an
// int returns ErrOverflow. A nil input (including a typed nil) returns
// ErrNilAny; every other type returns ErrUnsupported.
func ToIntStrict(a any) (int, error) {
	switch x := a.(type) {
	case int:
		return x, nil
	case int8:
		return int(x), nil
	case int16:
		return int(x), nil
	case int32:
		return int(x), nil
	case int64:
		return signedToInt(x)
	case uint:
		return unsignedToInt(uint64(x))
	case uint8:
		return int(x), nil
	case uint16:
		return int(x), nil
	case uint32:
		return unsignedToInt(uint64(x))
	case uint64:
		return unsignedToInt(x)
	case nil:
		return 0, ErrNilAny
	default:
		if arm.IsNil(a) {
			return 0, ErrNilAny
		}
		if rv := reflect.ValueOf(a); rv.Kind() == reflect.Ptr {
			return ToIntStrict(rv.Elem().Interface())
		}
		return 0, fmt.Errorf("%T %w", a, ErrUnsupported)
	}
}

// signedToInt narrows a signed 64-bit value to int, reporting ErrOverflow when
// it does not fit. On 64-bit platforms the bounds check is a no-op the compiler
// can elide; it only bites on 32-bit builds.
func signedToInt(v int64) (int, error) {
	if v < math.MinInt || v > math.MaxInt {
		return 0, fmt.Errorf("%d %w", v, ErrOverflow)
	}
	return int(v), nil
}

// unsignedToInt narrows an unsigned 64-bit value to int, reporting ErrOverflow
// when it exceeds the platform's int range.
func unsignedToInt(v uint64) (int, error) {
	if v > math.MaxInt {
		return 0, fmt.Errorf("%d %w", v, ErrOverflow)
	}
	return int(v), nil
}

// floatToInt truncates a float towards zero into an int, reporting ErrOverflow
// for NaN and values outside the int range. The high bound uses >= because the
// nearest float64 to math.MaxInt rounds up past it.
func floatToInt(v float64) (int, error) {
	if math.IsNaN(v) || v >= float64(math.MaxInt) || v < float64(math.MinInt) {
		return 0, fmt.Errorf("%v %w", v, ErrOverflow)
	}
	return int(v), nil
}
