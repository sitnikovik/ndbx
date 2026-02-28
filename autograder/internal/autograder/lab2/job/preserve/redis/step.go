package redis

import (
	"context"
	"time"
)

const (
	// Name is the name of the preserve Redis step.
	Name = "Preserve Redis Step"
	// Description is a brief explanation of what the preserve Redis step does.
	Description = "Checks if the user's session is preserved by verifying the TTL and session data in Redis."
)

// redisClient defines the interface for interacting with Redis.
type redisClient interface {
	// TTL retrieves the time-to-live of a key in Redis.
	TTL(
		ctx context.Context,
		key string,
	) (time.Duration, error)
	// HGetAll retrieves all fields and values of a hash stored at key.
	HGetAll(
		ctx context.Context,
		key string,
	) (map[string]string, error)
}

// Step represents the Redis preserve step in the autograder process.
type Step struct {
	// cli is the Redis client used to interact with the Redis server.
	cli redisClient
}

// NewStep creates a new Step instance with the provided Redis client.
func NewStep(cli redisClient) *Step {
	return &Step{
		cli: cli,
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
