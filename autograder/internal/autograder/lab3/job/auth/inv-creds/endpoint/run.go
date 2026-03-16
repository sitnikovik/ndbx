package endpoint

import (
	"bytes"
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	request "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/auth/login/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/resp"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/strings"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes the auth failed endpoint test step,
// sending a POST request with invalid credentials
// to the auth endpoint and validating the response.
func (s *Step) Run(
	_ context.Context,
	_ step.Variables,
) error {
	rsp, err := s.cli.PostJSON(
		endpoint.
			NewEndpoint(s.baseURL).
			Auth(),
		bytes.NewBuffer(
			request.
				NewBody(
					lab3.
						NewTestUser().
						Username(),
					"invalidpassword",
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
		response.AssertNotEmptyContent,
	)
	if err != nil {
		return errs.Wrap(
			err,
			"got unexpected response",
		)
	}
	body := resp.MustParseError(rsp.Body)
	err = strings.AssertEquals(
		"invalid credentials",
		body.Error(),
	)
	if err != nil {
		return errs.WrapJoin(
			"got unexpected error message",
			errs.ErrExpectationFailed,
			err,
		)
	}
	return nil
}
