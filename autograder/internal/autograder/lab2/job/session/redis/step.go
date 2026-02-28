package redis

import (
	"context"
	"time"
)

const (
	// Name is the name of the create session Redis step.
	Name = "Create Session Redis Step"
	// Description is a brief description of the create session Redis step.
	Description = "Checks the create session functionality by verifying that the session key in Redis" +
		" has been updated with a new TTL after a create session request."
)

// client defines the interface for interacting with Redis to perform the create session step.
type client interface {
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

// Step represents the Redis create session step.
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

// Description returns a brief description of what the step does.
func (s *Step) Description() string {
	return Description
}
