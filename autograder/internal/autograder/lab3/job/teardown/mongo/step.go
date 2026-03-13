package mongo

import "context"

const (
	// Name is the name of the MongoDB teardown step.
	Name = "MongoDB Teardown"
	// Description provides a brief explanation of what the MongoDB teardown step does.
	Description = "Tears down the MongoDB setup for Lab 3"
)

// mongoClient defines the interface for interacting with MongoDB to perform the teardown job for Lab 3.
type mongoClient interface {
	// Close closes the connection to the MongoDB server.
	Close(ctx context.Context) error
	// DropAll removes all data from the MongoDB server to ensure a clean state for Lab 3.
	DropAll(ctx context.Context) error
}

// Step represents the MongoDB teardown step for Lab 3 in the autograder process.
type Step struct {
	// mongo is the MongoDB client used to interact with the MongoDB server during the teardown process.
	mongo mongoClient
}

// NewStep creates a new Step instance with the provided MongoDB client.
func NewStep(mongo mongoClient) *Step {
	return &Step{
		mongo: mongo,
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
