package arm

import (
	"errors"
	"fmt"
	"runtime"
)

// Try runs given function in "protected mode",
// catching panics if they are thrown within it.
func Try(f func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = errors.Join(e, getStackTrace())
			} else {
				err = errors.Join(fmt.Errorf("%v", r), getStackTrace())
			}
		}
	}()

	err = f()
	return
}

// getStackTrace acquires current stacktrace and returns it
func getStackTrace() stacktrace {
	buf := make([]byte, 2048)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			return buf[:n]
		}
		buf = make([]byte, 2*len(buf))
	}
}

// stacktrace contains stack trace information and
// implements both Stringer and error interfaces.
type stacktrace []byte

func (e stacktrace) Error() string  { return string(e) }
func (e stacktrace) String() string { return string(e) }
