package ok

import (
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run searches the user by the endpoint and compares the result with the expected one.
func (s *Step) Run(
	_ context.Context,
	vars step.Variables,
) error {
	rsp, err := s.cli.Get(
		endpoint.
			NewEndpoint(s.baseURL).
			User(
					vars.
						MustGet(s.user.Hash()).
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
		return errs.Wrap(
			err,
			"got unexpected response",
		)
	}
	return nil
}
