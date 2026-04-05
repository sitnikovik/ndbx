package setup

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run truncates the keyspace in Apache Cassandra
// and returns an error if failed.
func (s *Step) Run(
	ctx context.Context,
	_ step.Variables,
) error {
	err := s.cli.TruncateKeyspace(ctx)
	if err != nil {
		return errs.WrapJoin(
			"failed to truncate",
			errs.ErrExternalDependencyFailed,
			err,
		)
	}
	return nil
}
