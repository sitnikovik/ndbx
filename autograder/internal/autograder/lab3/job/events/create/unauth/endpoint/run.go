package endpoint

import (
	"bytes"
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	request "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/post/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes the create event endpoint test step,
// sending a POST request to the create event endpoint and validating the response
// and checking that unauthenticated users cannot create events and receive the appropriate error response.
func (s *Step) Run(
	_ context.Context,
	_ step.Variables,
) error {
	rsp, err := s.cli.PostJSON(
		endpoint.
			NewEndpoint(s.baseURL).
			Events(),
		bytes.NewBuffer(
			request.
				NewBody(lab3.NewTestEvent()).
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
