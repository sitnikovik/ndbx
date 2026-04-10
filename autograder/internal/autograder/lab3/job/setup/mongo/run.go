package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes the MongoDB setup step for Lab 3.
//
// It pings the MongoDB service to ensure it is reachable and responsive,
// and drops all data from it to ensure a clean state for the autograder.
//
// Returns an error if any of the steps fail.
func (s *Step) Run(
	ctx context.Context,
	_ step.Variables,
) error {
	err := s.mongo.Ping(ctx)
	if err != nil {
		return errs.Wrap(err, "failed to ping MongoDB")
	}
	err = s.mongo.DropAll(ctx)
	if err != nil {
		return errs.Wrap(err, "failed to drop MongoDB data")
	}
	return nil
}
