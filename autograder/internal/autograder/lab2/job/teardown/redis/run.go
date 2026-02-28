package redis

import (
	"context"
	"fmt"

	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes the teardown operation for Lab 2.
//
// It flushes all data from the Redis server
// to clean state after all steps are done.
//
// Returns an error if any of the steps fail.
func (j *Job) Run(
	ctx context.Context,
	_ step.Variables,
) error {
	err := j.cli.FlushAll(ctx)
	if err != nil {
		return fmt.Errorf("failed to flush Redis: %w", err)
	}
	err = j.cli.Close(ctx)
	if err != nil {
		return fmt.Errorf("failed to close Redis client: %w", err)
	}
	return nil
}
