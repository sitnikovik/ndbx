package redis

import "context"

const (
	// Name is the name of the Redis setup step for Lab 3 in the autograder process.
	Name = "Redis Setup"
	// Description provides a brief description of the Redis setup step for Lab 3.
	Description = "Sets up the Redis environment for Lab 3."
)

// redisClient defines the interface for interacting with Redis to perform the setup job.
type redisClient interface {
	// Ping checks the connection to the Redis server.
	Ping(ctx context.Context) error
	// FlushAll removes all keys from the Redis server.
	FlushAll(ctx context.Context) error
}

// Step represents the Redis setup step for Lab 3 in the autograder process.
type Step struct {
	// redis is the Redis client used to interact with the Redis server during the setup process.
	redis redisClient
}

// NewStep creates a new Step instance with the provided Redis client.
func NewStep(redisClient redisClient) *Step {
	return &Step{
		redis: redisClient,
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
