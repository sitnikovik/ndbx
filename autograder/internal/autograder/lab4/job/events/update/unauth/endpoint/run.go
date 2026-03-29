package endpoint

import (
	"bytes"
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	request "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/patch/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run runs PATCH request to change the event
// and validates the response to be like unauthorized.
func (s *Step) Run(
	_ context.Context,
	vars step.Variables,
) error {
	rsp, err := s.cli.Patch(
		endpoint.
			NewEndpoint(s.baseURL).
			Event(
				variable.
					NewValues(vars).
					MustEventID(),
			),
		bytes.NewBuffer(
			request.
				NewBody(
					request.WithCategory("party"),
					request.WithCity("Москва"),
					request.WithPrice(1000),
				).
				MustBytes(),
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
		response.AssertUnauthorizedStatus,
		response.AssertEmptyContent,
	)
	if err != nil {
		return errs.WrapJoin(
			"got unexpected response",
			errs.ErrExpectationFailed,
			err,
		)
	}
	return nil
}
