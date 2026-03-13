package redis

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes the Redis teardown step for Lab 3.
//
// It flushes all data from the Redis service to ensure a clean state for the autograder,
// and closes the connection to it.
//
// Returns an error if any of the steps fail.
func (s *Step) Run(
	ctx context.Context,
	_ step.Variables,
) error {
	err := s.redis.FlushAll(ctx)
	if err != nil {
		return errs.Wrap(err, "failed to flush Redis")
	}
	err = s.redis.Close(ctx)
	if err != nil {
		return errs.Wrap(err, "failed to close Redis connection")
	}
	return nil
}
