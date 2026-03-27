package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

const (
	// Name is the name of the step.
	Name = "Create an event by Mongo"
	// Description is a brief description of the step.
	Description = "Creates the provived event by MongoDB request that to be found by fitler in the next steps"
)

// mongoClient defines the interface for interacting with MongoDB.
type mongoClient interface {
	// Insert inserts the list of documents into the specified collection.
	Insert(
		ctx context.Context,
		collection string,
		kvs ...doc.KVs,
	) error
}

// Step represents the HTTP create event step in the autograder process.
type Step struct {
	// mongo is the MongoDB client used to interact with the database.
	mongo mongoClient
	// event is the event that has to be created by the target application.
	event event.Event
}

// NewStep creates a new Step instance.
func NewStep(
	mongo mongoClient,
	event event.Event,
) *Step {
	return &Step{
		mongo: mongo,
		event: event,
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
