package redis

import (
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run performs the health check operation
// to verify that the Redis instance is empty.
func (s *Step) Run(
	ctx context.Context,
	_ step.Variables,
) error {
	empty, err := s.cli.Empty(ctx)
	if err != nil {
		return errors.Join(
			errs.ErrRedisFailed,
			err,
		)
	}
	if !empty {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"expected Redis to be empty, but it is not",
		)
	}
	return nil
}
