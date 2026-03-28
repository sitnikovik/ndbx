package notfound

import (
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run tries to get a non-existent user and validates the response.
func (s *Step) Run(
	_ context.Context,
	vars step.Variables,
) error {
	rsp, err := s.cli.Get(
		endpoint.
			NewEndpoint(s.baseURL).
			User(s.user.ID().String()),
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
		response.AssertNotFoundStatus,
		response.AssertEmptyContent,
	)
	if err != nil {
		return errs.Wrap(
			err,
			"got unexpected response",
		)
	}
	return nil
}
