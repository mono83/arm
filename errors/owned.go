package errors

import (
	"errors"
	"fmt"
)

// ErrNoCause is the placeholder cause substituted when a caused error is
// constructed with a nil cause.
var ErrNoCause = errors.New("no cause given")

// owned is an error carrying the name of its owner alongside a formattable
// message pattern and its arguments.
type owned struct {
	owner   string
	pattern string
	args    []any
}

// ErrorOwner returns the name of the error's owner.
func (o owned) ErrorOwner() string { return o.owner }

// ErrorArgs returns the message formatting arguments, if any.
func (o owned) ErrorArgs() []any { return o.args }

// Error returns the owner-prefixed, formatted message.
func (o owned) Error() string {
	if len(o.args) > 0 {
		return o.owner + ": " + fmt.Sprintf(o.pattern, o.args...)
	}
	return o.owner + ": " + o.pattern
}

// NewOwned builds an error owned by owner with a static message.
func NewOwned(owner, message string) error {
	return owned{owner: owner, pattern: message}
}

// NewOwnedf builds an error owned by owner with a printf-style message.
func NewOwnedf(owner, pattern string, args ...any) error {
	return owned{owner: owner, pattern: pattern, args: args}
}

// ownedCaused is an owned error that additionally wraps an underlying cause.
type ownedCaused struct {
	owned
	cause error
}

// Unwrap returns the underlying cause, enabling errors.Is and errors.As
// traversal.
func (o ownedCaused) Unwrap() error { return o.cause }

// Error returns the owner-prefixed, formatted message followed by the cause.
func (o ownedCaused) Error() string {
	return o.owned.Error() + ": " + o.cause.Error()
}

// NewOwnedCaused builds an error owned by owner with a static message wrapping
// cause. A nil cause is replaced with ErrNoCause.
func NewOwnedCaused(owner, message string, cause error) error {
	if cause == nil {
		cause = ErrNoCause
	}
	return ownedCaused{
		owned: owned{owner: owner, pattern: message},
		cause: cause,
	}
}

// NewOwnedCausedf builds an error owned by owner with a printf-style message
// wrapping cause. A nil cause is replaced with ErrNoCause.
func NewOwnedCausedf(owner, pattern string, cause error, args ...any) error {
	if cause == nil {
		cause = ErrNoCause
	}
	return ownedCaused{
		owned: owned{owner: owner, pattern: pattern, args: args},
		cause: cause,
	}
}
