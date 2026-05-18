package endpoint

import (
	"bytes"
	"context"
	"errors"
	"net/http"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/post/resp/body"
	request "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/post/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
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
			variable.EventID,
			body.
				MustParseBody(rsp.Body).
				ID(),
		)
	}
	return nil
}
