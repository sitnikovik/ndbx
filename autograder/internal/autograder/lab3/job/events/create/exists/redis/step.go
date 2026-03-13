package redis

import (
	"context"
)

const (
	// Name is the name of the step.
	Name = "Create Event Exists Redis Step"
	// Description is a brief description of the step.
	Description = "Checks if an event with the same ID already exists in Redis" +
		" and validates its fields to ensure it was created correctly"
)

// redisClient defines the interface for interacting with Redis.
type redisClient interface {
	// HGetAll retrieves all fields and values of a hash stored at key.
	HGetAll(
		ctx context.Context,
		key string,
	) (map[string]string, error)
}

// Step represents the Redis create event exists step in the autograder process.
type Step struct {
	// redis is the Redis client used to interact with the database.
	redis redisClient
}

// NewStep creates a new Step instance with the provided Redis client.
func NewStep(cli redisClient) *Step {
	return &Step{
		redis: cli,
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
