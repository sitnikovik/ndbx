package neo4j

import "github.com/sitnikovik/ndbx/autograder/internal/errs"

// Validate validates the Neo4j configuration.
func (c Config) Validate() error {
	err := c.conn.Validate()
	if err != nil {
		return errs.Wrap(err, "connection")
	}
	err = c.auth.Validate()
	if err != nil {
		return errs.Wrap(err, "auth")
	}
	return nil
}
