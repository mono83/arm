package arm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoe(t *testing.T) {
	if x, err := Todoe[int](); assert.Error(t, err) {
		assert.Equal(t, 0, x)
		assert.Equal(t, `accessing not implemented value of type "int"`, err.Error())
	}
	if x, err := Todoe[*int](); assert.Error(t, err) {
		assert.Nil(t, x)
		assert.Equal(t, `accessing not implemented value of type "*int"`, err.Error())
	}
	if x, err := Todoe[TestInterface](); assert.Error(t, err) {
		assert.Nil(t, x)
		assert.Equal(t, `accessing not implemented value of type "arm.TestInterface"`, err.Error())
	}
}

func TestTodoPanic(t *testing.T) {
	assert.Panics(t, func() {
		Todo[int]()
	})
}

type TestInterface interface{}
