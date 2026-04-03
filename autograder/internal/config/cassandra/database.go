package cassandra

import "github.com/sitnikovik/ndbx/autograder/internal/client/cassandra/consistency"

// Database is a Cassandra database configuration.
type Database struct {
	// keyspace is the Cassandra keyspace to use.
	keyspace string
	// consistency is the consistency level for Cassandra queries.
	consistency consistency.Consistency
}

// NewDatabase creates a new Database instance.
func NewDatabase(
	keyspace string,
	consistency consistency.Consistency,
) Database {
	return Database{
		keyspace:    keyspace,
		consistency: consistency,
	}
}

// Keyspace returns the keyspace to use for Cassandra queries.
func (d Database) Keyspace() string {
	return d.keyspace
}

// Consistency returns the consistency level for Cassandra queries.
func (d Database) Consistency() consistency.Consistency {
	return d.consistency
}
