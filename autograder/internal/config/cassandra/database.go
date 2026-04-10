package cassandra

import (
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/client/cassandra/consistency"
)

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

// Validate validates the connection config
// and returns if it has any invalid value.
func (d Database) Validate() error {
	if d.Keyspace() == "" {
		return errors.New("empty keyspace")
	}
	if d.Consistency().String() == "" {
		return errors.New("unknown consistency level")
	}
	return nil
}
