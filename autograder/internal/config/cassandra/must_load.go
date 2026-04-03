package cassandra

import (
	"github.com/sitnikovik/ndbx/autograder/internal/client/cassandra/consistency"
	"github.com/sitnikovik/ndbx/autograder/internal/env"
)

// MustLoad loads Cassandra configuration from environment variables
// and returns a new Config instance
// or oanics if any required variable is missing.
func MustLoad() Config {
	return NewConfig(
		NewConnection(
			env.
				MustGet("CASSANDRA_HOSTS").
				Strings(),
			env.
				MustGet("CASSANDRA_PORT").
				MustInt(),
		),
		NewAuth(
			env.
				Get("CASSANDRA_USERNAME").
				String(),
			env.
				Get("CASSANDRA_PASSWORD").
				String(),
		),
		NewDatabase(
			env.
				MustGet("CASSANDRA_KEYSPACE").
				String(),
			consistency.
				MustParseConsistency(
					env.
						MustGet("CASSANDRA_CONSISTENCY").
						String(),
				),
		),
	)
}
