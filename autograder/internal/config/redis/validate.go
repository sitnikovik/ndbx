package redis

import (
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// Validate checks if the Redis configuration is valid.
func (c Config) Validate() error {
	if c.addr == "" {
		return errs.Wrap(errs.ErrInvalidConfig, "redis address is required")
	}
	if c.db < 0 {
		return errs.Wrap(errs.ErrInvalidConfig, "invalid redis DB")
	}
	return nil
}
