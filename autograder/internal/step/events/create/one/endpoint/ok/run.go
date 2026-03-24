package ok

import (
	"bytes"
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/post/resp/body"
	request "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/post/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/strings"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run creates the event specified on the Step created and validates the response.
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
				NewBody(s.event).
				MustBytes(),
		),
	)
	if err != nil {
		return errors.Join(
			errs.ErrHTTPFailed,
			err,
		)
	}
	defer errs.MustBeClosed(
		rsp.Body.Close(),
	)
	err = response.AssertAll(
		rsp,
		response.AssertCreatedStatus,
		response.AssertNotEmptyContent,
	)
	if err != nil {
		return errs.WrapJoin(
			"got unexpected response",
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
	v := body.MustParseBody(rsp.Body)
	err = strings.AssertNotEmpty(v.ID())
	if err != nil {
		return errs.WrapNested(
			errs.ErrExpectationFailed,
			err,
			"got empty event ID in response body",
		)
	}
	return nil
}
