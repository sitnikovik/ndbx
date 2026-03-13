package redis

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes the Redis setup step for Lab 3.
//
// It pings the Redis service to ensure it is reachable and responsive,
// and flushes all data from it to ensure a clean state for the autograder.
//
// Returns an error if any of the steps fail.
func (s *Step) Run(
	ctx context.Context,
	_ step.Variables,
) error {
	err := s.redis.Ping(ctx)
	if err != nil {
		return errs.Wrap(err, "failed to ping Redis")
	}
	err = s.redis.FlushAll(ctx)
	if err != nil {
		return errs.Wrap(err, "failed to flush Redis")
	}
	return nil
}
