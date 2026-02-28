package redis

import (
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/cookie"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	"github.com/sitnikovik/ndbx/autograder/internal/user/session"
)

// Run executes the expectation.
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
	exists, err := s.cli.Has(ctx, key)
	if err != nil {
		return errs.Wrap(
			errors.Join(
				errs.ErrRedisFailed,
				err,
			),
			"failed to check existence of Redis key %s",
			log.String(key),
		)
	}
	if exists {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"expected Redis key %s to be expired, but it still exists",
			log.String(key),
		)
	}
	return nil
}
