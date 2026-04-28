package redis

import (
	"context"
	"errors"
	"strconv"

	"github.com/sitnikovik/ndbx/autograder/internal/app/redis/event/reviews"
	"github.com/sitnikovik/ndbx/autograder/internal/app/redis/event/reviews/field"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/times/duration"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run checks for the event reviews data
// meets expectations and validates TTL in Redis.
func (s *Step) Run(
	ctx context.Context,
	_ step.Variables,
) error {
	k := reviews.Key(
		s.event.
			Content().
			TitleHash(),
	)
	m, err := s.cli.HGetAll(ctx, k)
	if err != nil {
		return errors.Join(
			errs.ErrExternalDependencyFailed,
			err,
		)
	}
	rate, err := strconv.ParseFloat(m[field.Rating], 64)
	if err != nil {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"failed to parse '%s' value: %v",
			field.Rating,
			err,
		)
	}
	err = numbers.AssertEquals(
		s.expect.Counts().Rating().Exact(),
		rate,
	)
	if err != nil {
		return errs.Wrap(
			err,
			"got unexpected rating of reviews",
		)
	}
	count, err := strconv.Atoi(m[field.Count])
	if err != nil {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"failed to parse '%s' value: %v",
			field.Count,
			err,
		)
	}
	err = numbers.AssertEquals(
		s.expect.Counts().Count(),
		count,
	)
	if err != nil {
		return errs.Wrap(
			err,
			"got unexpected amount of reviews",
		)
	}
	ttl, err := s.cli.TTL(ctx, k)
	if err != nil {
		return errors.Join(
			errs.ErrExternalDependencyFailed,
			err,
		)
	}
	err = duration.AssertEquals(s.expect.TTL(), ttl)
	if err != nil {
		return errs.Wrap(
			err,
			"unexpected TTL",
		)
	}
	return nil
}
