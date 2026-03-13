package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes the MongoDB teardown step for Lab 3.
//
// It drops all data from the MongoDB service to ensure a clean state for the autograder,
// and closes the connection to it.
//
// Returns an error if any of the steps fail.
func (s *Step) Run(
	ctx context.Context,
	_ step.Variables,
) error {
	err := s.mongo.DropAll(ctx)
	if err != nil {
		return errs.Wrap(err, "failed to drop MongoDB data")
	}
	err = s.mongo.Close(ctx)
	if err != nil {
		return errs.Wrap(err, "failed to close MongoDB connection")
	}
	return nil
}
