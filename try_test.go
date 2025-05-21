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
	err := errors.New("some error")
	if e := Try(func() error { return err }); e != nil {
		assert.Equal(t, err, e)
	}
}

func TestTry_Panic(t *testing.T) {
	if err := Try(func() error { panic("some panic") }); assert.Error(t, err) {
		assert.False(t, errors.As(err, new(runtime.Error)))
		assert.Equal(t, err, errors.New("some panic"))
	}
}

func TestTry_Nil(t *testing.T) {
	if err := Try(nil); assert.Error(t, err) {
		assert.True(t, errors.As(err, new(runtime.Error)))
	}
}
