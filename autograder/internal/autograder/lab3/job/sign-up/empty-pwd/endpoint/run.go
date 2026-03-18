package endpoint

import (
	"bytes"
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/resp"
	request "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/users/post/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/strings"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes the sign-up endpoint test step,
// sending a POST request to the sign-up endpoint and validating the response.
func (s *Step) Run(
	_ context.Context,
	_ step.Variables,
) error {
	rsp, err := s.cli.PostJSON(
		endpoint.
			NewEndpoint(s.baseURL).
			SignUp(),
		bytes.NewBuffer(
			request.
				NewBody(
					lab3.NewTestUser(),
					"",
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
		response.AssertBadRequestStatus,
		response.AssertNotEmptyContent,
	)
	if err != nil {
		return errs.Wrap(
			err,
			"got unexpected response",
		)
	}
	body := resp.MustParseError(rsp.Body)
	err = strings.AssertNotEmpty(body.Error())
	if err != nil {
		return errs.WrapJoin(
			"response has no message",
			errs.ErrExpectationFailed,
			err,
		)
	}
	return nil
}
