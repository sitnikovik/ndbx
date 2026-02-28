package redis

import (
	"context"
	"fmt"

	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes the setup operation for Lab 2.
//
// It performs the following steps:
//  1. Pings the Redis server to ensure it is reachable and responsive.
//  2. Flushes all data from the Redis server to ensure a clean state for the autograder.
//
// It returns an error if any of the steps fail.
func (j *Job) Run(
	ctx context.Context,
	_ step.Variables,
) error {
	err := j.cli.Ping(ctx)
	if err != nil {
		return fmt.Errorf("failed to ping Redis: %w", err)
	}
	err = j.cli.FlushAll(ctx)
	if err != nil {
		return fmt.Errorf("failed to flush Redis: %w", err)
	}
	return nil
}
