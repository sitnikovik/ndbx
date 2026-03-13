package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

const (
	// Name is the name of the step.
	Name = "Create Event Mongo Step"
	// Description is a brief description of the step.
	Description = "Checks if the event document was created in the database after creating an event, " +
		"by querying the MongoDB collection and verifying that the expected document exists."
)

// mongoClient defines the interface for interacting with MongoDB.
type mongoClient interface {
	// AllBy retrieves all documents from the specified collection that match the given key-value pairs.
	AllBy(
		ctx context.Context,
		collection string,
		by doc.KVs,
	) (doc.Documents, error)
	// ByID retrieves a single document from the specified collection that matches the given ID.
	ByID(
		ctx context.Context,
		collection string,
		id string,
	) (doc.Document, error)
	// Indexes retrieves the list of indexes for the specified collection.
	Indexes(
		ctx context.Context,
		collection string,
	) (doc.Indexes, error)
}

// Step represents the MongoDB create event step in the autograder process.
type Step struct {
	// mongo is the MongoDB client used to interact with the database.
	mongo mongoClient
}

// NewStep creates a new Step instance with the provided MongoDB client.
func NewStep(cli mongoClient) *Step {
	return &Step{
		mongo: cli,
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
