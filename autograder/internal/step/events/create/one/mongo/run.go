package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/event/collection"
	appdoc "github.com/sitnikovik/ndbx/autograder/internal/app/mongo/event/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run inserts an event into MongoDB
// and sets the event id got into the step variables.
func (s *Step) Run(
	ctx context.Context,
	vars step.Variables,
) error {
	ids, err := s.mongo.Insert(
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
	if len(ids) == 0 {
		return errs.Wrap(
			errs.ErrExternalDependencyFailed,
			"got empty ids after insert",
		)
	}
	vars.Set(s.event.Hash(), ids[0])
	return nil
}
