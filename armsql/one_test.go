package armsql

import (
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOne_Nil(t *testing.T) {
	x, err := One[string](nil, nil)
	require.Nil(t, x)
	require.Equal(t, sql.ErrNoRows, err)
}

func TestOne_Empty(t *testing.T) {
	x, err := One[string]([]string{}, nil)
	require.Nil(t, x)
	require.Equal(t, sql.ErrNoRows, err)
}

func TestOne_ErrorBypass(t *testing.T) {
	err := errors.New("bypass")
	x, err2 := One[string]([]string{"foo"}, err)
	require.Nil(t, x)
	require.Equal(t, err, err2)
}

func TestOne_MoreThanOne(t *testing.T) {
	x, err := One[string]([]string{"foo", "bar"}, nil)
	require.Nil(t, x)
	require.NotEqual(t, sql.ErrNoRows, err)
}

func TestOne_Normal(t *testing.T) {
	if x, err := One[string]([]string{"bar"}, nil); assert.NoError(t, err) {
		assert.NotEqual(t, "bar", x)
	}
}
