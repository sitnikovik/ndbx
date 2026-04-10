package endpoint

import (
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes the like endpoint test step,
// sending a POST request to the like endpoint and validating the response.
func (s *Step) Run(
	_ context.Context,
	vars step.Variables,
) error {
	rsp, err := s.cli.PostJSON(
		endpoint.
			NewEndpoint(s.baseURL).
			EventLike(
				vars.
					MustGet(s.event.Hash()).
					AsString(),
			),
		nil,
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
		response.AssertNoContentStatus,
		response.AssertEmptyContent,
	)
	if err != nil {
		return errs.Wrap(
			err,
			"got unexpected response",
		)
	}
	cksess := session.MustParseSession(
		rsp.Cookies(),
	)
	err = cksess.Validate()
	if err != nil {
		return errs.Wrap(
			err,
			"invalid session value in cookie",
		)
	}
	err = cksess.MatchVariables(vars)
	if err != nil {
		return errs.Wrap(
			err,
			"session cookie does not match expected variables",
		)
	}
	return nil
}
