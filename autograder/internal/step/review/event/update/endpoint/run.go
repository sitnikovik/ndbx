package endpoint

import (
	"bytes"
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run updates the event review and validates the response.
func (s *Step) Run(
	_ context.Context,
	vars step.Variables,
) error {
	rsp, err := s.cli.Patch(
		endpoint.
			NewEndpoint(s.baseURL).
			EventReview(
				vars.
					MustGet(s.event.Hash()).
					AsString(),
				vars.
					MustGet(variable.Review+s.event.Hash()).
					AsString(),
			),
		bytes.NewBuffer(
			s.rq.MustBytes(),
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
	defer func() {
		errs.MustBeClosed(
			rsp.Body.Close(),
		)
	}()
	err = s.want.Assert(rsp)
	if err != nil {
		return errs.WrapJoin(
			"unexpected response",
			errs.ErrExpectationFailed,
			err,
		)
	}
	return nil
}
