package config

import (
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// Validate checks if the job configuration is valid
// and returns an error if any part of the configuration is invalid.
func (c Config) Validate() error {
	err := c.redis.Validate()
	if err != nil {
		return errs.WrapNested(errs.ErrInvalidConfig, err, "redis")
	}
	if err := c.mongo.Validate(); err != nil {
		return errs.WrapNested(errs.ErrInvalidConfig, err, "mongo")
	}
	err = c.app.Validate()
	if err != nil {
		return errs.WrapNested(errs.ErrInvalidConfig, err, "app")
	}
	return nil
}
