package arm

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestTry_Normal(t *testing.T) {
	assert.NoError(t, Try(func() error {
		return nil
	}))
}

func TestTry_Error(t *testing.T) {
	given := errors.New("some error")
	if err := Try(func() error { return given }); err != nil {
		assert.Equal(t, given, err)
		assert.False(t, errors.As(err, new(stacktrace)))
	}
}

func TestTry_Panic(t *testing.T) {
	if err := Try(func() error { panic("some panic") }); assert.Error(t, err) {
		assert.False(t, errors.As(err, new(runtime.Error)))
		assert.True(t, errors.As(err, new(stacktrace)))
	}
}

func TestTry_Nil(t *testing.T) {
	if err := Try(nil); assert.Error(t, err) {
		assert.True(t, errors.As(err, new(runtime.Error)))
		assert.True(t, errors.As(err, new(stacktrace)))
	}
}
