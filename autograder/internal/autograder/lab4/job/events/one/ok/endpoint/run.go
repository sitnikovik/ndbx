package endpoint

import (
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/one/resp/body"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run retrieves the event by the endpont
// and compares it with the event stored in the vars.
func (s *Step) Run(
	_ context.Context,
	vars step.Variables,
) error {
	vals := variable.NewValues(vars)
	ev := vals.MustEvent()
	rsp, err := s.cli.Get(
		endpoint.
			NewEndpoint(s.baseURL).
			Event(ev.ID().String()),
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
		return errs.Wrap(err, "got unexpected response")
	}
	err = expect.AssertEquals(
		ev,
		body.
			MustParseBody(rsp.Body).
			Event(),
	)
	if err != nil {
		return errs.Wrap(err, "got unexpected event")
	}
	return nil
}
