package redis

import (
	"context"
	"encoding/json"

	"github.com/sitnikovik/ndbx/autograder/internal/app/reaction/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/redis/user/recommendations"
	"github.com/sitnikovik/ndbx/autograder/internal/app/redis/user/recommendations/field"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/times/duration"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run checks that the user recommendations are correct.
func (s *Step) Run(
	ctx context.Context,
	vars step.Variables,
) error {
	k := recommendations.Key(
		vars.
			MustGet(s.user.Hash()).
			AsString(),
	)
	m, err := s.cli.HGetAll(ctx, k)
	if err != nil {
		return errs.WrapJoin(
			"failed to get recommendations",
			errs.ErrExternalDependencyFailed,
			err,
		)
	}
	jsn, ok := m[field.Events]
	if !ok {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"field '%s' not found",
			field.Events,
		)
	}
	if s.want.HasEvents() {
		var got []event.Event
		err = json.Unmarshal([]byte(jsn), &got)
		if err != nil {
			return errs.WrapJoin(
				"failed to unmarshal events",
				errs.ErrMarshallFailed,
				errs.ErrExpectationFailed,
				err,
			)
		}
		err = numbers.AssertEquals(
			len(s.want.Events()),
			len(got),
		)
		if err != nil {
			return errs.Wrap(
				err,
				"unexpected count of events",
			)
		}
	}
	ttl, err := s.cli.TTL(ctx, k)
	if err != nil {
		return errs.WrapJoin(
			"failed to get ttl",
			errs.ErrExternalDependencyFailed,
			err,
		)
	}
	if s.want.HasTTL() {
		err = duration.AssertEquals(ttl, s.want.TTL())
		if err != nil {
			return errs.Wrap(
				err,
				"unexpected TTL",
			)
		}
	}
	return nil
}
