package setup

import "context"

const (
	// Name is the name of the step.
	Name = "Cassandra Setup"
	// Description provides a brief description of the step.
	Description = "Sets up the Apache Cassandra environment for tests."
)

// cassandraClient defines the interface for interacting
// with Apache Cassandra to perform the setup job.
type cassandraClient interface {
	// TruncateKeyspace truncates all tables in the keyspace.
	TruncateKeyspace(ctx context.Context) error
}

// Step represents the Apache Cassandra setup step for tests.
type Step struct {
	// cli is the Apache Cassandra client used to perform the setup job.
	cli cassandraClient
}

// NewStep creates a new Step instance
// with the provided Apache Cassandra client.
func NewStep(cli cassandraClient) *Step {
	return &Step{
		cli: cli,
	}
}

// Name returns the name of the step.
func (s *Step) Name() string {
	return Name
}

// Description returns a brief explanation of what the step does.
func (s *Step) Description() string {
	return Description
}
