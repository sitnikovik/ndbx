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
	if createdAt.UTC().After(time.Now().UTC()) {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"session creation time is in the future: %s",
			log.String(tim),
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
	vars.Set(redis.SessionCreatedAtField, createdAt)
	vars.Set(redis.SessionTTL, ttl)
	return nil
}
