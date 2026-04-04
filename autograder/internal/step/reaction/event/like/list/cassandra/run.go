package cassandra

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra/event/reaction"
	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra/event/reaction/filter"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	qb "github.com/sitnikovik/ndbx/autograder/internal/client/cassandra/query/builder"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run selects likes by filter related to the event
// and validates the number of likes got.
func (s *Step) Run(
	ctx context.Context,
	vars step.Variables,
) error {
	likes, err := reaction.
		NewLikes(
			s.cassandra,
			reaction.WithFilter(
				filter.NewFilter(
					qb.NewWhere(),
					filter.WithEventID(
						event.ID(
							vars.
								MustGet(s.event.Hash()).
								AsString(),
						),
					),
					filter.WithCreatedBy(
						user.NewID(
							vars.
								MustGet(s.event.Created().By().Hash()).
								AsString(),
						),
					),
				),
			),
			reaction.WithLimit(s.expected),
		).
		Select(ctx)
	if err != nil {
		return errs.WrapJoin(
			"failed to select likes",
			errs.ErrExternalDependencyFailed,
			err,
		)
	}
	err = numbers.AssertEquals(
		s.expected,
		len(likes),
	)
	if err != nil {
		return errs.Wrap(
			err,
			"unexpected number of likes",
		)
	}
	return nil
}
