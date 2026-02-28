package session

import (
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// Validate checks if the Value configuration is valid
// and returns an error if it is not.
func (c Config) Validate() error {
	if c.ttl <= 0 {
		return errs.Wrap(
			errs.ErrInvalidConfig,
			"TTL must be greater than zero",
		)
	}
	return nil
}
