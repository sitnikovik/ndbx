package user

import (
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// Validate checks the validity of the Config fields
// and returns an error if any field is invalid.
func (c Config) Validate() error {
	err := c.session.Validate()
	if err != nil {
		return errs.WrapNested(
			errs.ErrInvalidConfig,
			err,
			"session",
		)
	}
	return nil
}
