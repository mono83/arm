package anies

import "reflect"

// basic unwraps a named scalar type (e.g. `type UserID int`) into a value of
// its predeclared base type: int64, uint64, float32, float64, bool or string.
// Converters recurse on the result to treat the alias like the builtin it is
// defined over.
//
// It reports ok=false for predeclared types (empty PkgPath) and non-scalar
// kinds: predeclared scalars already match an explicit case in each converter's
// type switch, so returning false both stops unbounded recursion and leaves
// structs, slices and the like unsupported. float32 is kept as float32 so
// ToString preserves its shorter 32-bit representation.
func basic(rv reflect.Value) (any, bool) {
	if rv.Type().PkgPath() == "" {
		return nil, false
	}
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int(), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return rv.Uint(), true
	case reflect.Float32:
		return float32(rv.Float()), true
	case reflect.Float64:
		return rv.Float(), true
	case reflect.Bool:
		return rv.Bool(), true
	case reflect.String:
		return rv.String(), true
	}
	return nil, false
}
