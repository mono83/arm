package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewOwned(t *testing.T) {
	err := NewOwned("svc", "boom")
	require.EqualError(t, err, "svc: boom")

	o := err.(owned)
	require.Equal(t, "svc", o.ErrorOwner())
	require.Nil(t, o.ErrorArgs())
}

func TestNewOwnedf(t *testing.T) {
	err := NewOwnedf("svc", "boom %d/%s", 42, "x")
	require.EqualError(t, err, "svc: boom 42/x")

	o := err.(owned)
	require.Equal(t, "svc", o.ErrorOwner())
	require.Equal(t, []any{42, "x"}, o.ErrorArgs())
}

func TestNewOwnedCaused(t *testing.T) {
	cause := errors.New("root")
	err := NewOwnedCaused("svc", "boom", cause)
	require.EqualError(t, err, "svc: boom: root")
	require.ErrorIs(t, err, cause)

	o := err.(ownedCaused)
	require.Equal(t, "svc", o.ErrorOwner())
	require.Nil(t, o.ErrorArgs())
}

func TestNewOwnedCausedf(t *testing.T) {
	cause := errors.New("root")
	err := NewOwnedCausedf("svc", "boom %d", cause, 42)
	require.EqualError(t, err, "svc: boom 42: root")
	require.ErrorIs(t, err, cause)
	require.Equal(t, []any{42}, err.(ownedCaused).ErrorArgs())
}

func TestNewOwnedCausedNilCause(t *testing.T) {
	t.Run("plain", func(t *testing.T) {
		err := NewOwnedCaused("svc", "boom", nil)
		require.EqualError(t, err, "svc: boom: no cause given")
		require.ErrorIs(t, err, ErrNoCause)
	})
	t.Run("formatted", func(t *testing.T) {
		err := NewOwnedCausedf("svc", "boom %d", nil, 42)
		require.EqualError(t, err, "svc: boom 42: no cause given")
		require.ErrorIs(t, err, ErrNoCause)
	})
}
