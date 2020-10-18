package permissions

import "errors"

var (
	// ErrNoAllower is returned when Allow is called on a nil Allower.
	ErrNoAllower = errors.New("cannot call Allow on a nil Allower")
)
