package errs

import "errors"

// WrapJoin wraps multiple errors into a single error with the custom message.
func WrapJoin(msg string, err ...error) error {
	return Wrap(errors.Join(err...), msg)
}
