package cassandra

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
)

const (
	// Name is the name of the step.
	Name = "Get reactions for the event"
	// Description is a brief description of the step.
	Description = "Retrieves the event reactions of compares it with the provided ones"
)

// cassandraClient implements the client
// to interact with Apache Cassandra.
type cassandraClient interface {
	// Select selects the rows from the given query
	// and returns an iterator to scan the rows.
	Select(
		ctx context.Context,
		query string,
		args ...any,
	) (cassandra.Scanner, error)
}

// Step represents the step
// to get the event reactions from Apache Cassandra and compare and validate them.
type Step struct {
	// cassandra is the Cassandra client to run queries.
	cassandra cassandraClient
	// event is the event releated to the likes to search.
	event event.Event
	// expected is the number of likes to expect.
	expected int
}

// NewStep creates a new Step instance.
func NewStep(
	cli cassandraClient,
	evnt event.Event,
	expected int,
) *Step {
	return &Step{
		cassandra: cli,
		event:     evnt,
		expected:  expected,
	}
}

// Name returns the name of the step.
func (s *Step) Name() string {
	return Name
}

// Description returns a brief description of what the step does.
func (s *Step) Description() string {
	return Description
}
