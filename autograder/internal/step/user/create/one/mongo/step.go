package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
)

const (
	// Name is the name of the step.
	Name = "Create a user by Mongo"
	// Description is a brief description of the step.
	Description = "Creates the provived user by MongoDB request that to be found by fitler in the next steps"
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

// Step represents the HTTP create user step in the autograder process.
type Step struct {
	// mongo is the MongoDB client used to interact with the database.
	mongo mongoClient
	// user is the user that has to be created by the target application.
	user user.User
}

// NewStep creates a new Step instance.
func NewStep(
	mongo mongoClient,
	user user.User,
) *Step {
	return &Step{
		mongo: mongo,
		user: user,
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
