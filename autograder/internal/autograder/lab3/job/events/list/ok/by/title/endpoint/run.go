package endpoint

import (
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/resp/body"
	rq "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes the create event endpoint test step,
// sending a POST request to the create event endpoint and validating the response.
func (s *Step) Run(
	_ context.Context,
	_ step.Variables,
) error {
	rsp, err := s.cli.Get(
		endpoint.WithQuery(
			endpoint.
				NewEndpoint(s.baseURL).
				Events(),
			rq.
				NewBody(
					rq.WithTitle("Test Event"),
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
	err = numbers.AssertEquals(0, len(body.Events()))
	if err != nil {
		return errs.WrapNested(
			errs.ErrExpectationFailed,
			err,
			"expected exactly 1 event in response",
		)
	}
	err = numbers.AssertEquals(0, body.Count())
	if err != nil {
		return errs.WrapNested(
			errs.ErrExpectationFailed,
			err,
			"expected count field to be equal to events count",
		)
	}
	return nil
}
