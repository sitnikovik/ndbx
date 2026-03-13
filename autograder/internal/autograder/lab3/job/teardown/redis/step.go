package redis

import "context"

const (
	// Name is the name of the Redis teardown job.
	Name = "Redis Teardown"
	// Description provides a brief explanation of what the Redis teardown job does.
	Description = "Tears down the Redis environment for Lab 3"
)

// redisClient defines the interface for interacting with Redis to perform the teardown job.
type redisClient interface {
	// Close closes the connection to the Redis server.
	Close(ctx context.Context) error
	// FlushAll removes all keys from the Redis server.
	FlushAll(ctx context.Context) error
}

// Step represents the Redis teardown step for Lab 3 in the autograder process.
type Step struct {
	// rediscli is the Redis client used to interact with the Redis server during the teardown process.
	redis redisClient
}

// NewStep creates a new Step instance with the provided Redis client.
func NewStep(redis redisClient) *Step {
	return &Step{
		redis: redis,
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
