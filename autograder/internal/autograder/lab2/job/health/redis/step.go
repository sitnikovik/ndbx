package redis

import (
	"context"
)

const (
	// Name is the name of the health check job.
	Name = "Healthcheck Redis step"
	// Description is a brief description of the health check job.
	Description = "Checks the health of the Redis server by verifying that it is empty before the test runs."
)

// redisClient defines the interface for interacting with Redis to perform the health check.
type redisClient interface {
	Empty(ctx context.Context) (bool, error)
}

// Step represents the Redis health check step in the autograder process.
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

// Description returns a brief description of what the step does.
func (s *Step) Description() string {
	return Description
}
