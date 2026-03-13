package redis

import (
	"context"
)

const (
	// Name is the name of the logout Redis step.
	Name = "Auth Logout Redis Step"
	// Description is a brief explanation of what the logout Redis step does.
	Description = "Checks if the user's session has been invalidated" +
		" by verifying the absence of the session key in Redis after logout."
)

// client defines the interface for interacting with Redis to perform the logout step.
type client interface {
	// Has checks if the key exists in Redis.
	Has(
		ctx context.Context,
		key string,
	) (bool, error)
}

// Step represents the Redis logout step.
type Step struct {
	// cli is the Redis client used to interact with the Redis server.
	cli client
}

// NewStep creates a new logout Step instance with the provided Redis client.
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
