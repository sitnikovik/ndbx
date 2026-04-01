package redis

import (
	"context"
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
)

const (
	// Name is the name of the step.
	Name = "Likes in Redis"
	// Description is a brief description of the step.
	Description = "Verifies that a like has been stored in Redis after calling the endpoint"
)

// redisClient defines the interface
// to interact the Redis service to perform the session refresh step.
type redisClient interface {
	// TTL retrieves the time-to-live of a key in Redis.
	TTL(
		ctx context.Context,
		key string,
	) (time.Duration, error)
	// HGet retrieves the value associated
	// with a field in a hash stored at key.
	HGet(
		ctx context.Context,
		key string,
		field string,
	) (string, error)
}

// Step represents the HTTP create event step in the autograder process.
type Step struct {
	// cli is Redis client used to interact with the Redis server.
	cli redisClient
	// event is the event that has been liked.
	event event.Event
	// expect is the amount of likes to expect.
	expect int
	// ttl is the time-to-live of the like in Redis.
	ttl time.Duration
}

// NewStep creates a new Step instance
// with the provided Redis client,
// the event to like, number of	likes to expect and TTL of the like.
func NewStep(
	cli redisClient,
	evnt event.Event,
	expect int,
	ttl time.Duration,
) *Step {
	return &Step{
		cli:    cli,
		event:  evnt,
		expect: expect,
		ttl:    ttl,
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
