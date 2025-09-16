package arm

import "reflect"

// Todo is a temporary stubber function to use for places when some
// value/component is expected but yet not implemented. Any invocation
// of this function will produce panic.
func Todo[a any]() a {
	_, err := Todoe[a]()
	panic(err)
}

// Todoe is a temporary stubber function to use for places when some
// value/component is expected but yet not implemented. Any invocation
// of this function will produce an error response.
func Todoe[a any]() (val a, err error) {
	t := reflect.TypeOf(new(a))
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	err = ErrTodo(`accessing not implemented value of type "` + t.String() + `"`)
	return
}

// ErrTodo is an error returned on TODO assertions
type ErrTodo string

func (e ErrTodo) Error() string { return string(e) }
