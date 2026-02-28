package redis

import (
	"context"
)

const (
	// Name is the name of the expire Redis step.
	Name = "Expire Redis Step"
	// Description is a brief explanation of what the expire Redis step does.
	Description = "Checks if the user's session has expired by verifying the existence of the session key in Redis."
)

// client defines the interface for interacting with Redis to perform the expire step.
type client interface {
	// Has checks if the key exists in Redis.
	Has(
		ctx context.Context,
		key string,
	) (bool, error)
}

// Step represents the Redis expire step.
type Step struct {
	// cli is the Redis client used to interact with the Redis server.
	cli client
}

// NewStep creates a new Step instance with the provided Redis client.
func NewStep(cli client) *Step {
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
