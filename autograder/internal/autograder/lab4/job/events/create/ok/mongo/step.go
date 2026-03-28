package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/shard"
)

const (
	// Name is the name of the step.
	Name = "Bulk Event Creation"
	// Description is a brief description of the step.
	Description = "Runs bulk event creation and checks how it stored in MongoDB"
)

// mongoClient defines the interface for interacting with MongoDB.
type mongoClient interface {
	// Insert inserts the list of documents into the specified collection.
	// Returns a slice of inserted document IDs and an error if any.
	Insert(
		ctx context.Context,
		collection string,
		kvs ...doc.KVs,
	) ([]string, error)
	// Shards retrieves the shard
	// information for the specified collection.
	Shards(
		ctx context.Context,
		collection string,
	) (shard.Shards, error)
	// HostsOfShard returns a list of hosts where the shard is running.
	HostsOfShard(
		ctx context.Context,
		id string,
	) ([]string, error)
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
