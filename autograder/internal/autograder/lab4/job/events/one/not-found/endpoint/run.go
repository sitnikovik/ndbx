package endpoint

import (
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run tries to get non-existent event
// and validates the response got by the endpoint.
func (s *Step) Run(
	_ context.Context,
	_ step.Variables,
) error {
	rsp, err := s.cli.Get(
		endpoint.
			NewEndpoint(s.baseURL).
			Event("2ip3ue32-9ejnojkdsdp932u8eji"),
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
		response.AssertNotEmptyContent,
	)
	if err != nil {
		return errs.Wrap(err, "got unexpected response")
	}
	return nil
}
