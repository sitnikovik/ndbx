package cassandra

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra/event/review"
	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra/event/review/filter"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	qb "github.com/sitnikovik/ndbx/autograder/internal/client/cassandra/query/builder"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run selects reviews by filter related to the event
// and validates the number of reviews got.
func (s *Step) Run(
	ctx context.Context,
	vars step.Variables,
) error {
	list, err := review.
		NewReviews(
			s.cassandra,
			review.WithFilter(
				filter.NewFilter(
					qb.NewWhere(),
					filter.WithEventID(
						event.ID(
							vars.
								MustGet(s.event.Hash()).
								AsString(),
						),
					),
				),
			),
			review.WithLimit(s.want.Count()+10),
		).
		Select(ctx)
	if err != nil {
		return errs.WrapJoin(
			"failed to get reviews",
			errs.ErrExternalDependencyFailed,
			err,
		)
	}
	err = numbers.AssertEquals(
		s.want.Count(),
		len(list),
	)
	if err != nil {
		return errs.Wrap(
			err,
			"unexpected number of reviews",
		)
	}
	return nil
}
