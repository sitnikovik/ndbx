package redis

import (
	"context"
	"errors"
	"strconv"

	"github.com/sitnikovik/ndbx/autograder/internal/app/redis/event/reactions"
	"github.com/sitnikovik/ndbx/autograder/internal/app/redis/event/reactions/field"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/times/duration"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run checks for the number of dislikes
// meets expectations and validates TTL in Redis.
func (s *Step) Run(
	ctx context.Context,
	_ step.Variables,
) error {
	k := reactions.Key(
		s.event.
			Content().
			TitleHash(),
	)
	v, err := s.cli.HGet(ctx, k, field.Dislikes)
	if err != nil {
		return errors.Join(
			errs.ErrExternalDependencyFailed,
			err,
		)
	}
	dislikes, err := strconv.Atoi(v)
	if err != nil {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"failed to parse 'dislikes' value: %v",
			err,
		)
	}
	err = numbers.AssertEquals(
		s.expect,
		dislikes,
	)
	if dislikes != s.expect {
		return errs.Wrap(
			err,
			"got unexpected amout of dislikes",
		)
	}
	ttl, err := s.cli.TTL(ctx, k)
	if err != nil {
		return errors.Join(
			errs.ErrExternalDependencyFailed,
			err,
		)
	}
	err = duration.AssertEquals(s.ttl, ttl)
	if err != nil {
		return errs.Wrap(
			err,
			"unexpected TTL",
		)
	}
	return nil
}
