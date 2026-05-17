package redis

import (
	"context"
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	"github.com/sitnikovik/ndbx/autograder/internal/step/user/one/recommendations/redis/expect"
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
	// want holds the expectations we need to check in the Step.
	want expect.Expectations
	// user is the user's which events are expected to be retrieved.
	user user.User
	// desc is the description of the step.
	desc step.Desc
}

// NewStep creates a new Step instance.
func NewStep(
	desc step.Desc,
	cli redisClient,
	usr user.User,
	want expect.Expectations,
) *Step {
	s := &Step{
		desc: desc,
		cli:  cli,
		user: usr,
		want: want,
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
