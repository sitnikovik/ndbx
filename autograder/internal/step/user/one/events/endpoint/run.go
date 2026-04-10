package endpoint

import (
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/resp/body"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes the search of user's events by filters and validates the response got.
func (s *Step) Run(
	_ context.Context,
	vars step.Variables,
) error {
	rsp, err := s.cli.Get(
		endpoint.WithQuery(
			endpoint.
				NewEndpoint(s.baseURL).
				UserEvents(
					vars.
						MustGet(s.user.Hash()).
						AsString(),
				),
			s.rq.URLQuery(),
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
	body := body.MustParseBody(rsp.Body)
	events := body.Events()
	err = numbers.AssertEquals(
		len(s.events),
		len(events),
	)
	if err != nil {
		return errs.Wrap(
			err,
			"got unexpected count of events",
		)
	}
	err = numbers.AssertEquals(
		len(s.events),
		body.Count(),
	)
	if err != nil {
		return errs.Wrap(
			err,
			"got unexpected 'count' field",
		)
	}
	if s.expect.HasReactions() {
		want := s.expect.Reactions()
		err = numbers.AssertEquals(
			len(want),
			len(events),
		)
		if err != nil {
			return errs.Wrap(
				err,
				"got length mismatch of expected reactions with gotten events",
			)
		}
		for i, ev := range events {
			err = expect.AssertEquals(
				want[i],
				ev.Reactions(),
			)
			if err != nil {
				return errs.Wrap(
					err,
					"got unexpected reactions for event '%s'",
					ev.ID().String(),
				)
			}
		}
	}
	return nil
}
