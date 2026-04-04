package config

import (
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// Validate checks if the job configuration is valid
// and returns an error if any part of the configuration is invalid.
func (c Config) Validate() error {
	err := c.redis.Validate()
	if err != nil {
		return errs.WrapJoin(
			"redis",
			errs.ErrInvalidConfig,
			err,
		)
	}
	err = c.mongo.Validate()
	if err != nil {
		return errs.WrapJoin(
			"mongo",
			errs.ErrInvalidConfig,
			err,
		)
	}
	err = c.cassandra.Validate()
	if err != nil {
		return errs.WrapJoin(
			"cassandra",
			errs.ErrInvalidConfig,
			err,
		)
	}
	err = c.app.Validate()
	if err != nil {
		return errs.WrapJoin(
			"app",
			errs.ErrInvalidConfig,
			err,
		)
	}
	return nil
}
