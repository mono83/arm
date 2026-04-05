package arm

import "errors"

// AllOfProvided2 calls each provider in order, returning early on the first
// nil provider or the first error. On success all return values are populated.
func AllOfProvided2[T1, T2 any](f1 func() (T1, error), f2 func() (T2, error)) (t1 T1, t2 T2, err error) {
	if f1 == nil {
		err = errors.New("f1 is nil")
		return
	}
	if f2 == nil {
		err = errors.New("f2 is nil")
		return
	}

	t1, err = f1()
	if err != nil {
		return
	}
	t2, err = f2()
	return
}

// AllOfProvided3 calls each provider in order, returning early on the first
// nil provider or the first error. On success all return values are populated.
func AllOfProvided3[T1, T2, T3 any](f1 func() (T1, error), f2 func() (T2, error), f3 func() (T3, error)) (t1 T1, t2 T2, t3 T3, err error) {
	if f1 == nil {
		err = errors.New("f1 is nil")
		return
	}
	if f2 == nil {
		err = errors.New("f2 is nil")
		return
	}
	if f3 == nil {
		err = errors.New("f3 is nil")
		return
	}

	t1, err = f1()
	if err != nil {
		return
	}
	t2, err = f2()
	if err != nil {
		return
	}
	t3, err = f3()
	return
}

// AllOfProvided4 calls each provider in order, returning early on the first
// nil provider or the first error. On success all return values are populated.
func AllOfProvided4[T1, T2, T3, T4 any](f1 func() (T1, error), f2 func() (T2, error), f3 func() (T3, error), f4 func() (T4, error)) (t1 T1, t2 T2, t3 T3, t4 T4, err error) {
	if f1 == nil {
		err = errors.New("f1 is nil")
		return
	}
	if f2 == nil {
		err = errors.New("f2 is nil")
		return
	}
	if f3 == nil {
		err = errors.New("f3 is nil")
		return
	}
	if f4 == nil {
		err = errors.New("f4 is nil")
		return
	}

	t1, err = f1()
	if err != nil {
		return
	}
	t2, err = f2()
	if err != nil {
		return
	}
	t3, err = f3()
	if err != nil {
		return
	}
	t4, err = f4()
	return
}

// AllOfProvided5 calls each provider in order, returning early on the first
// nil provider or the first error. On success all return values are populated.
func AllOfProvided5[T1, T2, T3, T4, T5 any](f1 func() (T1, error), f2 func() (T2, error), f3 func() (T3, error), f4 func() (T4, error), f5 func() (T5, error)) (t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, err error) {
	if f1 == nil {
		err = errors.New("f1 is nil")
		return
	}
	if f2 == nil {
		err = errors.New("f2 is nil")
		return
	}
	if f3 == nil {
		err = errors.New("f3 is nil")
		return
	}
	if f4 == nil {
		err = errors.New("f4 is nil")
		return
	}
	if f5 == nil {
		err = errors.New("f5 is nil")
		return
	}

	t1, err = f1()
	if err != nil {
		return
	}
	t2, err = f2()
	if err != nil {
		return
	}
	t3, err = f3()
	if err != nil {
		return
	}
	t4, err = f4()
	if err != nil {
		return
	}
	t5, err = f5()
	return
}
