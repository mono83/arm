package validate

import "errors"

// All validates every given value and returns the joined errors of those that
// fail, or nil when all are valid. Unlike fail-fast validation, it checks every
// value rather than stopping at the first error.
func All(vv ...Validable) error {
	var errs []error
	for _, v := range vv {
		if err := v.Validate(); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}
