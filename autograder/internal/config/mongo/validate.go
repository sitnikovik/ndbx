package mongo

import "github.com/sitnikovik/ndbx/autograder/internal/errs"

// Validate checks if the MongoDB configuration is valid
// and returns an error if any required field is missing or invalid.
func (c Config) Validate() error {
	if c.db == "" {
		return errs.Wrap(errs.ErrInvalidConfig, "database name is required")
	}
	if c.host == "" {
		return errs.Wrap(errs.ErrInvalidConfig, "host is required")
	}
	if c.port <= 0 || c.port > 99999 {
		return errs.Wrap(errs.ErrInvalidConfig, "invalid port")
	}
	return nil
}
