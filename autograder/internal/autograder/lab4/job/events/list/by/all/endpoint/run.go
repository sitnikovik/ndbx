package endpoint

import (
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/resp/body"
	rq "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/category"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

// Run executes the search of events by filters and valudates the response got.
func (s *Step) Run(
	_ context.Context,
	vars step.Variables,
) error {
	rsp, err := s.cli.Get(
		endpoint.WithQuery(
			endpoint.
				NewEndpoint(s.baseURL).
				Events(),
			rq.
				NewBody(
					rq.WithTitle("чудес"),
					rq.WithCategory(category.Exhibition),
					rq.WithEntryPrice(0, 0),
					rq.WithDates(
						timex.MustRFC3339("2026-03-24T00:00:00Z"),
						timex.MustRFC3339("2026-03-24T00:00:00Z"),
					),
					rq.WithByUser(
						user.NewIdentity(
							user.NewID("123"),
						),
					),
					rq.WithAddress("Ходынский"),
					rq.WithCity("Москва"),
				).
				URLQuery(),
		),
	)
	if err != nil {
		return errors.Join(
			errs.ErrHTTPFailed,
			err,
		)
	}
	defer func() {
		errs.MustBeClosed(
			rsp.Body.Close(),
		)
	}()
	err = response.AssertAll(
		rsp,
		response.AssertOKStatus,
		response.AssertNotEmptyContent,
	)
	if err != nil {
		return errs.WrapJoin(
			"got unexpected response",
			errs.ErrExpectationFailed,
			err,
		)
	}
	body := body.MustParseBody(rsp.Body)
	err = numbers.AssertEquals(1, len(body.Events()))
	if err != nil {
		return errs.WrapNested(
			errs.ErrExpectationFailed,
			err,
			"expected exactly 1 event in response",
		)
	}
	vars.Set(
		variable.EventID,
		body.
			Events()[0].
			ID().
			String(),
	)
	return nil
}
