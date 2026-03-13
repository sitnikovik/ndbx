package redis

import (
	"context"
)

const (
	// Name is the name of the step.
	Name = "Sign-Up Redis Step"
	// Description is a brief description of the step.
	Description = "Check if the user session created in Redis after signing up" +
		" by querying the Redis database and verifying that the expected session key exists."
)

// redisClient defines the interface for interacting with Redis.
type redisClient interface {
	// HGetAll retrieves all fields and values of a hash stored at key.
	HGetAll(
		ctx context.Context,
		key string,
	) (map[string]string, error)
}

// Step represents the Redis sign-up step in the autograder process.
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
