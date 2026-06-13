package arm

import "reflect"

// IsNil reports whether the given value is nil.
//
// Unlike a plain `a == nil` comparison, IsNil also detects typed nil values
// stored inside an interface. A nil pointer, map, slice, channel, function or
// interface assigned to an `any` produces a non-nil interface holding a nil
// dynamic value, for which `a == nil` evaluates to false. IsNil unwraps such
// values via reflection and returns true.
//
// For non-nilable kinds (numbers, strings, structs, arrays, booleans) IsNil
// always returns false.
func IsNil(a any) bool {
	if a == nil {
		return true
	}

	switch v := reflect.ValueOf(a); v.Kind() {
	case reflect.Ptr,
		reflect.Map,
		reflect.Slice,
		reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.UnsafePointer:
		return v.IsNil()
	default:
		return false
	}
}
