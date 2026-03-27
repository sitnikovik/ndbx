package ok

import (
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/users/list/resp/body"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes the search of events by filters and valudates the response got.
func (s *Step) Run(
	_ context.Context,
	_ step.Variables,
) error {
	rsp, err := s.cli.Get(
		endpoint.WithQuery(
			endpoint.
				NewEndpoint(s.baseURL).
				Events(),
			s.rq.URLQuery(),
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
		response.AssertOKStatus,
		response.AssertNotEmptyContent,
	)
	if err != nil {
		return errs.Wrap(
			err,
			"got unexpected response",
		)
	}
	body := body.MustParseBody(rsp.Body)
	err = numbers.AssertEquals(
		len(s.users),
		len(body.Users()),
	)
	if err != nil {
		return errs.Wrap(
			err,
			"got unexpected count of users",
		)
	}
	err = numbers.AssertEquals(
		len(s.users),
		body.Count(),
	)
	if err != nil {
		return errs.Wrap(
			err,
			"got unexpected 'count' field",
		)
	}
	return nil
}
