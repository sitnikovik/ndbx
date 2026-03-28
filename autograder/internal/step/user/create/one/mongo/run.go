package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/user/collection"
	appdoc "github.com/sitnikovik/ndbx/autograder/internal/app/mongo/user/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run inserts a user into MongoDB
// and sets the retrieved id into the step variables.
func (s *Step) Run(
	ctx context.Context,
	vars step.Variables,
) error {
	ids, err := s.mongo.Insert(
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
	if len(ids) == 0 {
		return errs.Wrap(
			errs.ErrExternalDependencyFailed,
			"got empty ids after insert",
		)
	}
	vars.Set(s.user.Hash(), ids[0])
	return nil
}
