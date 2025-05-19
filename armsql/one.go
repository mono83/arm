package armsql

import (
	"database/sql"
	"fmt"
)

// One returns exactly one element of given slice and produces
// next errors:
// 1. Given error if it is not nil
// 2. sql.ErrNoRows if empty slice given
// 3. Custom error if slice contains more than one element
func One[T any](slice []T, err error) (*T, error) {
	if err != nil {
		return nil, err
	}
	if ln := len(slice); ln == 0 {
		return nil, sql.ErrNoRows
	} else if ln > 1 {
		// This should not happen at all
		// Encountering this situation means either incorrect SQL query
		// or some critical database inconsistency.
		return nil, fmt.Errorf("expected only 1 row, got %d", ln)
	}

	return &slice[0], nil
}
