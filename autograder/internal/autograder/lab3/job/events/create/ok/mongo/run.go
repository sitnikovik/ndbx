package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/event/collection"
	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/event/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/event/doc/key"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/strings"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run checks that the event document created
// in the previous step matches the expected test event
// and that the required indexes exist in MongoDB.
func (s *Step) Run(
	ctx context.Context,
	vars step.Variables,
) error {
	vals := variable.NewValues(vars)
	evdoc, err := s.mongo.ByID(
		ctx,
		collection.Name,
		vals.MustEventID(),
	)
	if err != nil {
		return errs.WrapJoin(
			"failed to get events from MongoDB",
			errs.ErrExternalDependencyFailed,
			err,
		)
	}
	evnt := doc.NewEventDocument(evdoc).ToEvent()
	err = strings.AssertEquals(
		vals.
			MustUser().
			ID().
			String(),
		evnt.
			Created().
			By().
			ID().
			String(),
	)
	if err != nil {
		return errs.Wrap(
			err,
			"event creator ID does not match the expected user ID",
		)
	}
	idxx, err := s.mongo.Indexes(ctx, collection.Name)
	if err != nil {
		return errs.WrapJoin(
			"failed to get indexes from MongoDB",
			errs.ErrExternalDependencyFailed,
			err,
		)
	}
	idx := idxx.For(key.Title)
	if idx.Empty() {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"expected an index on the "+
				log.String(key.Title)+
				" field, but it does not exist",
		)
	}
	idx = idxx.For(key.CreatedBy)
	if idx.Empty() {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"expected an index on the "+
				log.String(key.CreatedBy)+
				" field, but it does not exist",
		)
	}
	idx = idxx.For(key.Title, key.CreatedBy)
	if idx.Empty() {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"expected an index on the "+
				log.String(key.Title)+
				" and "+
				log.String(key.CreatedBy)+
				" fields, but it does not exist",
		)
	}
	return nil
}
