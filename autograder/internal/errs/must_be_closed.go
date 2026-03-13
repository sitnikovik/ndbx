package errs

import "errors"

// MustBeClosed panics if the provided error is not nil,
// indicating that closing a resource failed.
func MustBeClosed(err error) {
	if err != nil {
		panic(errors.Join(ErrCloseFailed, err))
	}
}
