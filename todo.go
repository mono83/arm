package arm

import "fmt"

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
	err = fmt.Errorf(`accessing not implemented value of type "%T"`, val)
	return
}
