package redis

import (
	"context"
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	"github.com/sitnikovik/ndbx/autograder/internal/step/events/one/review/redis/expect"
)

// redisClient defines the interface
// to interact the Redis service to perform the session refresh step.
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

// Step represents the HTTP create event step in the autograder process.
type Step struct {
	// cli is Redis client used to interact with the Redis server.
	cli redisClient
	// event is the event that has been liked.
	event event.Event
	// expect is the amount of likes to expect.
	expect expect.Expectations
	// desc is the description of the step.
	desc step.Desc
}

// NewStep creates a new Step instance
// with the provided Redis client,
// the event to like, number of	likes to expect and TTL of the like.
func NewStep(
	desc step.Desc,
	cli redisClient,
	evnt event.Event,
	opts ...Option,
) *Step {
	s := &Step{
		cli:   cli,
		event: evnt,
		desc:  desc,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Name returns the name of the step.
func (s *Step) Name() string {
	return s.desc.Title()
}

// Description returns a brief description of what the step does.
func (s *Step) Description() string {
	return s.desc.Description()
}
