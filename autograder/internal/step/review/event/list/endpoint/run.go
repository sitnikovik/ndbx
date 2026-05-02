package endpoint

import (
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/reviews/events/list/resp/body"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes the search of reviews for the event and validates the response got.
func (s *Step) Run(
	_ context.Context,
	vars step.Variables,
) error {
	rsp, err := s.cli.Get(
		endpoint.
			NewEndpoint(s.baseURL).
			EventReviews(
				vars.
					MustGet(s.event.Hash()).
					AsString(),
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
	b := body.MustParseBody(rsp.Body)
	reviews := b.Reviews()
	n := len(reviews)
	err = numbers.AssertEquals(s.expect.Count(), n)
	if err != nil {
		return errs.WrapJoin(
			"unexpected number of reviews",
			errs.ErrExpectationFailed,
			err,
		)
	}
	err = numbers.AssertEquals(n, b.Count())
	if err != nil {
		return errs.WrapJoin(
			"got mismatch in 'count' field",
			errs.ErrExpectationFailed,
			err,
		)
	}
	return nil
}
