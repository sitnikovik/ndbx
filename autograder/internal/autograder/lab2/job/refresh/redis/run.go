package redis

import (
	"context"
	"errors"
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/cookie"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/times"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/times/duration"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	"github.com/sitnikovik/ndbx/autograder/internal/user/session"
)

// Run performs the Redis session refresh step by retrieving session information from Redis
// and validating the session creation time and TTL.
func (s *Step) Run(
	ctx context.Context,
	vars step.Variables,
) error {
	sid := vars.
		MustGet(cookie.SessionName).
		AsString()
	err := session.Validate(sid)
	if err != nil {
		return errs.Wrap(
			errors.Join(
				errs.ErrExpectationFailed,
				err,
			),
			"invalid session id",
		)
	}
	key := redis.SessionKey(sid)
	m, err := s.cli.HGetAll(ctx, key)
	if err != nil {
		return errs.Wrap(
			errors.Join(
				errs.ErrRedisFailed,
				err,
			),
			"failed to get value by %s",
			log.String(key),
		)
	}
	tim, ok := m[redis.SessionCreatedAtField]
	if !ok {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"%s not found in hash",
			log.String(redis.SessionCreatedAtField),
		)
	}
	createdAt, err := time.Parse(time.RFC3339, tim)
	if err != nil {
		return errs.Wrap(
			errors.Join(
				errs.ErrExpectationFailed,
				err,
			),
			"invalid time format for %s",
			log.String(redis.SessionCreatedAtField),
		)
	}
	if createdAt.IsZero() {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"%s field is empty",
			log.String(redis.SessionCreatedAtField),
		)
	}
	err = times.AssertEquals(
		vars.
			MustGet(redis.SessionCreatedAtField).
			AsTime(),
		createdAt,
	)
	if err != nil {
		return errs.Wrap(
			err,
			"session creation seems to be updated",
		)
	}
	ttl, err := s.cli.TTL(ctx, key)
	if err != nil {
		return errs.Wrap(
			errors.Join(
				errs.ErrRedisFailed,
				err,
			),
			"failed to get TTL for %s",
			log.String(key),
		)
	}
	err = duration.AssertEquals(
		vars.
			MustGet(redis.SessionTTL).
			AsDuration(),
		ttl,
	)
	if err != nil {
		return errs.Wrap(
			err,
			"session TTL seems to be not updated",
		)
	}
	return nil
}
