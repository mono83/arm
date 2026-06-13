package anies

import "errors"

var (
	ErrNilAny      = errors.New("nil any")
	ErrUnsupported = errors.New("unsupported")
	ErrOverflow    = errors.New("overflow")
)
