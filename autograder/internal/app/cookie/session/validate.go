package session

import (
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
	"github.com/sitnikovik/ndbx/autograder/internal/user/session"
)

// Validate checks if the session cookie is valid and has the expected flags set
// and returns an error if any of the checks fail.
func (s Session) Validate() error {
	name := s.ck.Name
	err := session.Validate(s.ck.Value)
	if err != nil {
		return errors.Join(
			errs.ErrExpectationFailed,
			err,
		)
	}
	if !s.ck.HttpOnly {
		return errs.Wrap(
			errs.ErrMissedCookie,
			"expect %s cookie to have http only flag set",
			log.String(name),
		)
	}
	if s.Expired() {
		return errs.Wrap(
			errs.ErrMissedCookie,
			"expect %s cookie to have MaxAge flag set",
			log.String(name),
		)
	}
	return nil
}
