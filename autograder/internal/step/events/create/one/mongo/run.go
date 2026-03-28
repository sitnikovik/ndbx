package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/event/collection"
	appdoc "github.com/sitnikovik/ndbx/autograder/internal/app/mongo/event/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run inserts an event into MongoDB.
func (s *Step) Run(
	ctx context.Context,
	_ step.Variables,
) error {
	_, err := s.mongo.Insert(
		ctx,
		collection.Name,
		appdoc.FromEvent(s.event).KVs(),
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
