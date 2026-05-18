package setup

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run clears the Neo4j database
// and returns an error if failed.
func (s *Step) Run(
	ctx context.Context,
	_ step.Variables,
) error {
	err := s.cli.DeleteAll(ctx)
	if err != nil {
		return errs.WrapJoin(
			"failed to delete all data",
			errs.ErrExternalDependencyFailed,
			err,
		)
	}
	return nil
}