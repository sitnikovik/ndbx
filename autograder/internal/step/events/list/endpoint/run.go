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
		return errs.WrapJoin(
			"got unexpected response",
			errs.ErrExpectationFailed,
			err,
		)
	}
	body := body.MustParseBody(rsp.Body)
	events := body.Events()
	n := len(events)
	if s.expect.EventsRequired() {
		wantN := len(s.expect.Events())
		err = numbers.AssertEquals(
			wantN,
			n,
		)
		if err != nil {
			return errs.Wrap(
				err,
				"expected exactly %d event in response",
				len(s.expect.Events()),
			)
		}
		err = numbers.AssertEquals(
			wantN,
			body.Count(),
		)
		if err != nil {
			return errs.Wrap(
				err,
				"got mismatch in 'count' field",
			)
		}
	}
	if s.expect.ReactionsRequired() {
		want := s.expect.Reactions()
		err = numbers.AssertEquals(
			len(want),
			n,
		)
		if err != nil {
			panic("got length mismatch of expected reactions with gotten events")
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
	if s.expect.ReviewsRequired() {
		want := s.expect.Reviews()
		err = numbers.AssertEquals(
			len(want),
			n,
		)
		if err != nil {
			panic("got length mismatch of expected reviews with gotten events")
		}
		for i, ev := range events {
			err = expect.AssertEquals(
				want[i],
				ev.Reviews(),
			)
			if err != nil {
				return errs.Wrap(
					err,
					"got unexpected reviews for event '%s' %v != %v",
					ev.ID().String(),
					want[i],
					ev.Reviews(),
				)
			}
		}
	}
	return nil
}
