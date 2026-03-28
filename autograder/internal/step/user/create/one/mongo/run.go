package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/user/collection"
	appdoc "github.com/sitnikovik/ndbx/autograder/internal/app/mongo/user/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run inserts a user into MongoDB.
func (s *Step) Run(
	ctx context.Context,
	_ step.Variables,
) error {
	_, err := s.mongo.Insert(
		ctx,
		collection.Name,
		appdoc.FromUser(s.user).KVs(),
	)
	if err != nil {
		return errs.WrapJoin(
			"failed to insert",
			err,
			errs.ErrExternalDependencyFailed,
		)
	}
	return nil
}
