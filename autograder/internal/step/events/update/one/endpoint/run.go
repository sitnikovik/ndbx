package endpoint

import (
	"bytes"
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	request "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/post/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run creates the event specified on the Step created and validates the response.
func (s *Step) Run(
	_ context.Context,
	vars step.Variables,
) error {
	rsp, err := s.cli.Patch(
		endpoint.
			WithQuery(
				endpoint.
					NewEndpoint(s.baseURL).
					Event(
						vars.
							MustGet(s.event.Hash()).
							AsString(),
					),
				s.rq.URLQuery(),
			),
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
