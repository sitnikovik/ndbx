package cassandra

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	"github.com/sitnikovik/ndbx/autograder/internal/step/review/event/list/cassandra/expectation"
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
	// want is the expectations to compare the results with.
	want expectation.Expectations
	// desc is the description of the step.
	desc step.Desc
}

// NewStep creates a new Step instance.
func NewStep(
	desc step.Desc,
	cli cassandraClient,
	evnt event.Event,
	want expectation.Expectations,
) *Step {
	return &Step{
		desc:      desc,
		cassandra: cli,
		event:     evnt,
		want:      want,
	}
}

// Name returns the name of the step.
func (s *Step) Name() string {
	return s.desc.Title()
}

// Description returns a brief description of what the step does.
func (s *Step) Description() string {
	return s.desc.Description()
}
