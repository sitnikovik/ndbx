package mongo

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/event/collection"
	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/event/doc/key"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3"
	mongodoc "github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run checks that an event with the same ID already exists in MongoDB
// and there is only one such event.
func (s *Step) Run(
	ctx context.Context,
	_ step.Variables,
) error {
	events, err := s.mongo.AllBy(
		ctx,
		collection.Name,
		mongodoc.NewKVs(
			mongodoc.NewKV(
				key.Title,
				lab3.
					NewTestEvent().
					Content().
					Title(),
			),
		),
	)
	if err != nil {
		return errs.WrapJoin(
			"failed to get events from MongoDB",
			errs.ErrExternalDependencyFailed,
			err,
		)
	}
	err = numbers.AssertEquals(1, len(events))
	if err != nil {
		return errs.WrapJoin(
			"expected exactly one event",
			errs.ErrExpectationFailed,
			err,
		)
	}
	return nil
}
