package cassandra

import "github.com/sitnikovik/ndbx/autograder/internal/errs"

// Validate checks if the Apache Cassandra configuration is valid
// and returns an error if any required field is missing or invalid.
func (c Config) Validate() error {
	err := c.Connection().Validate()
	if err != nil {
		return errs.Wrap(err, "connection")
	}
	err = c.Database().Validate()
	if err != nil {
		return errs.Wrap(err, "database")
	}
	return nil
}
