package neo4j

import (
	"github.com/sitnikovik/ndbx/autograder/internal/env"
)

// MustLoad loads Neo4j configuration from environment variables
// and returns a new Config instance
// or panics if any required variable is missing.
func MustLoad() Config {
	return NewConfig(
		NewConnection(
			env.
				MustGet("NEO4J_URL").
				String(),
		),
		NewAuth(
			env.
				Get("NEO4J_USERNAME").
				String(),
			env.
				Get("NEO4J_PASSWORD").
				String(),
		),
	)
}
