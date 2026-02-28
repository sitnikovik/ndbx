package app

import (
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// Validate checks if the Config is valid and returns an error if it is not.
func (c Config) Validate() error {
	err := c.user.Validate()
	if err != nil {
		return errs.WrapNested(errs.ErrInvalidConfig, err, "user")
	}
	if c.host == "" {
		return errs.Wrap(errs.ErrInvalidConfig, "host is required")
	}
	if c.port <= 0 || c.port > 99999 {
		return errs.Wrap(errs.ErrInvalidConfig, "invalid port number")
	}
	return nil
}
