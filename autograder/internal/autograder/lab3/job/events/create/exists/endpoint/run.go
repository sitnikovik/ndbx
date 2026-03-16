package endpoint

import (
	"bytes"
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	request "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/post/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/resp"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/strings"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes the create event endpoint test step,
// sending a POST request to the create event endpoint and validating the response.
func (s *Step) Run(
	_ context.Context,
	vars step.Variables,
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
		response.AssertConflictStatus,
		response.AssertNotEmptyContent,
	)
	if err != nil {
		return errs.WrapJoin(
			"got unexpected response",
			errs.ErrExpectationFailed,
			err,
		)
	}
	body := resp.MustParseError(rsp.Body)
	err = strings.AssertNotEmpty(body.Error())
	if err != nil {
		return errs.WrapJoin(
			"response does not have message",
			errs.ErrExpectationFailed,
			err,
		)
	}
	cksess := session.MustParseSession(rsp.Cookies())
	err = cksess.Validate()
	if err != nil {
		return errs.WrapNested(
			errs.ErrExpectationFailed,
			err,
			"invalid session value in cookie",
		)
	}
	err = cksess.MatchVariables(vars)
	if err != nil {
		return errs.WrapNested(
			errs.ErrExpectationFailed,
			err,
			"session cookie does not match expected variables",
		)
	}
	return nil
}
