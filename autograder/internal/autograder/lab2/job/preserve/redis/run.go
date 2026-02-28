package redis

import (
	"context"
	"errors"
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/cookie"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	"github.com/sitnikovik/ndbx/autograder/internal/user/session"
)

// Run performs the Redis check-in step by retrieving session information from Redis
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
		return errors.Join(errs.ErrExpectationFailed, err)
	}
	key := redis.SessionKey(sid)
	m, err := s.cli.HGetAll(ctx, key)
	if err != nil {
		return errors.Join(
			errs.ErrRedisFailed,
			err,
		)
	}
	tim, ok := m[redis.SessionCreatedAtField]
	if !ok {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"%s field not found in hash",
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
	createdAtExpected := vars.
		MustGet(redis.SessionCreatedAtField).
		AsTime()
	if !createdAt.UTC().Equal(createdAtExpected) {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"time mismatch: expected %s, got %s",
			log.String(createdAtExpected.String()),
			log.String(createdAt.String()),
		)
	}
	ttl, err := s.cli.TTL(ctx, key)
	if err != nil {
		return errors.Join(
			errs.ErrRedisFailed,
			err,
		)
	}
	ttlExpected := vars.
		MustGet(redis.SessionTTL).
		AsDuration()
	expiresAtExpected := createdAt.Add(ttlExpected)
	expiresAt := createdAt.Add(ttl)
	timeDiff := expiresAt.Sub(expiresAtExpected)
	if timeDiff < 0 {
		timeDiff = -timeDiff
	}
	if timeDiff > 3*time.Second {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"expiration time mismatch: expected %s, got %s (diff: %s)",
			log.String(expiresAtExpected.String()),
			log.String(expiresAt.String()),
			log.String(timeDiff.String()),
		)
	}
	return nil
}
