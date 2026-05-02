package endpoint

import (
	"bytes"
	"context"
	"errors"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/reviews/events/create/resp/body"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run creates a review for the event and validates the response.
func (s *Step) Run(
	_ context.Context,
	vars step.Variables,
) error {
	rsp, err := s.cli.PostJSON(
		endpoint.
			NewEndpoint(s.baseURL).
			EventReviews(
				vars.
					MustGet(s.event.Hash()).
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
	if rsp.StatusCode == http.StatusCreated {
		vars.Set(
			variable.Review+s.event.Hash(),
			body.
				MustParseBody(rsp.Body).
				ID(),
		)
	}
	return nil
}
