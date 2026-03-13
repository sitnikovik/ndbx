package mongo

import "context"

const (
	// Name is the name of the MongoDB setup step for Lab 3 in the autograder process.
	Name = "Mongo Setup"
	// Description provides a brief description of the MongoDB setup step for Lab 3.
	Description = "Sets up the MongoDB environment for Lab 3"
)

// mongoClient defines the interface for interacting with MongoDB to perform the setup job for Lab 3.
type mongoClient interface {
	// Ping checks the connection to the MongoDB server.
	Ping(ctx context.Context) error
	// DropAll removes all data from the MongoDB server to ensure a clean state for Lab 3.
	DropAll(ctx context.Context) error
}

// Step represents the MongoDB setup step for Lab 3 in the autograder process.
type Step struct {
	// mongo is the MongoDB client used to interact with the MongoDB server during the setup process.
	mongo mongoClient
}

// NewStep creates a new Step instance with the provided MongoDB client.
func NewStep(mongoClient mongoClient) *Step {
	return &Step{
		mongo: mongoClient,
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
