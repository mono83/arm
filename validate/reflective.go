package validate

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/mono83/arm"
	armerrors "github.com/mono83/arm/errors"
)

var (
	// ErrNil is returned when a nil value is validated and nil values are not
	// allowed.
	ErrNil = errors.New("nil")
)

// ReflectiveValidator validates arbitrary values using reflection.
//
// Values implementing Validable are validated through their own Validate
// method. Pointers are dereferenced. Composite values (slices, arrays, maps
// and structs) are traversed element by element, each element being validated
// recursively. All other kinds are considered valid.
//
// The Validable check is a type assertion on the value as passed, so a type
// whose Validate method has a pointer receiver is only recognized when passed
// as a pointer. Passing such a type by value bypasses its Validate method and
// falls through to reflective field traversal.
//
// Cyclic structures are handled: a reference (pointer, slice or map) already
// present on the current traversal path is treated as valid rather than
// recursed into, so validation terminates instead of overflowing the stack.
type ReflectiveValidator struct {
	// AllowNil reports whether a nil value passes validation instead of
	// yielding ErrNil.
	AllowNil bool

	// SkipStruct disables recursion into struct fields.
	SkipStruct bool
	// SkipSliceValues disables recursion into slice and array elements.
	SkipSliceValues bool
	// SkipMapKeys disables recursion into map keys.
	SkipMapKeys bool
	// SkipMapValues disables recursion into map values.
	SkipMapValues bool
}

// Validate validates the given value, recursing into composite kinds. It
// returns the joined errors of every failing element, or nil when the value
// is valid.
func (r ReflectiveValidator) Validate(a any) error {
	return r.validate(a, visited{})
}

// validate performs the recursive traversal, carrying the set of references
// already entered on the current path to break cycles.
func (r ReflectiveValidator) validate(a any, seen visited) error {
	if arm.IsNil(a) {
		return arm.If(r.AllowNil, nil, ErrNil)
	}

	if x, ok := a.(Validable); ok {
		return x.Validate()
	}

	value := reflect.ValueOf(a)
	switch value.Kind() {
	case reflect.Pointer:
		cycle, leave := seen.enter(value)
		if cycle {
			return nil
		}
		defer leave()
		return r.validate(value.Elem().Interface(), seen)
	case reflect.Slice:
		if r.SkipSliceValues {
			return nil
		}
		cycle, leave := seen.enter(value)
		if cycle {
			return nil
		}
		defer leave()
		return r.validateElements(value, seen)
	case reflect.Array:
		if r.SkipSliceValues {
			return nil
		}
		return r.validateElements(value, seen)
	case reflect.Map:
		cycle, leave := seen.enter(value)
		if cycle {
			return nil
		}
		defer leave()
		return r.validateMap(value, seen)
	case reflect.Struct:
		if r.SkipStruct {
			return nil
		}
		return r.validateStruct(value, seen)
	default:
		return nil
	}
}

// validateElements validates every element of a slice or array, tagging each
// failure with its index.
func (r ReflectiveValidator) validateElements(value reflect.Value, seen visited) error {
	var errs []error
	for i := 0; i < value.Len(); i++ {
		err := at(fmt.Sprintf("[%d]", i), r.validate(value.Index(i).Interface(), seen))
		errs = appendErr(errs, err)
	}
	return errors.Join(errs...)
}

// validateMap validates map keys and values according to the SkipMapKeys and
// SkipMapValues flags, tagging each failure with its key.
func (r ReflectiveValidator) validateMap(value reflect.Value, seen visited) error {
	var errs []error
	for iter := value.MapRange(); iter.Next(); {
		key := iter.Key().Interface()
		if !r.SkipMapKeys {
			err := at(fmt.Sprintf("key[%v]", key), r.validate(key, seen))
			errs = appendErr(errs, err)
		}
		if !r.SkipMapValues {
			err := at(fmt.Sprintf("[%v]", key), r.validate(iter.Value().Interface(), seen))
			errs = appendErr(errs, err)
		}
	}
	return errors.Join(errs...)
}

// validateStruct validates every exported field of a struct, tagging each
// failure with its field name. Unexported fields are skipped, since their
// values cannot be read through reflection.
func (r ReflectiveValidator) validateStruct(value reflect.Value, seen visited) error {
	var errs []error
	typ := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if !field.CanInterface() {
			continue
		}
		err := at(typ.Field(i).Name, r.validate(field.Interface(), seen))
		errs = appendErr(errs, err)
	}
	return errors.Join(errs...)
}

// at wraps a non-nil error into a validate-owned error tagged with the
// location at which it occurred, keeping the original error unwrappable via
// errors.Is and errors.As. A nil error is returned unchanged.
//
// The location is passed as the message rather than a format pattern, so it is
// safe even when it contains percent signs (e.g. a map key like "50%").
func at(location string, err error) error {
	if err == nil {
		return nil
	}
	return armerrors.NewOwnedCausedf("validate", location, err)
}

// appendErr appends err to errs only when err is non-nil.
func appendErr(errs []error, err error) []error {
	if err == nil {
		return errs
	}
	return append(errs, err)
}

// cycleKey identifies a reference value by its dynamic type, underlying
// pointer and length. The length distinguishes overlapping sub-slices that
// share a backing array, so only a genuine self-reference collides.
type cycleKey struct {
	typ reflect.Type
	ptr uintptr
	len int
}

// visited is the set of references entered on the current traversal path.
type visited map[cycleKey]struct{}

// enter records the reference held by value as being on the current path. It
// reports whether the reference was already present (a cycle) and, when it was
// not, returns a leave function that must be deferred to drop it once its
// subtree has been traversed. Removing on exit keeps the set scoped to the
// active path, so a reference shared between sibling branches is validated in
// each rather than mistaken for a cycle.
func (v visited) enter(value reflect.Value) (cycle bool, leave func()) {
	key := cycleKey{typ: value.Type(), ptr: value.Pointer()}
	if value.Kind() == reflect.Slice {
		key.len = value.Len()
	}
	if _, ok := v[key]; ok {
		return true, func() {}
	}
	v[key] = struct{}{}
	return false, func() { delete(v, key) }
}
