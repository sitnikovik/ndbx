package redis

import (
	"context"

	sessionCookies "github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	"github.com/sitnikovik/ndbx/autograder/internal/app/redis/user/session"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/strings"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run checks if the session with the given session ID exists in Redis
// and retrieves its creation and update timestamps to set them as variables
// for further steps in the autograder process.
func (s *Step) Run(
	ctx context.Context,
	vars step.Variables,
) error {
	sid := vars.MustGet(sessionCookies.Name).AsString()
	m, err := s.redis.HGetAll(ctx, session.Key(sid))
	if err != nil {
		return errs.WrapJoin(
			"failed to get session",
			errs.ErrExternalDependencyFailed,
			err,
		)
	}
	val := session.NewValue(sid, m)
	err = val.Validate()
	if err != nil {
		return errs.WrapNested(
			err,
			errs.ErrExpectationFailed,
			"invalid session data",
		)
	}
	sess := val.MustToSession()
	err = strings.AssertEquals(
		variable.
			NewValues(vars).
			MustUser().
			ID().
			String(),
		sess.
			User().
			ID().
			String(),
	)
	if err != nil {
		return errs.Wrap(
			err,
			"session user ID mismatch",
		)
	}
	if !sess.Updated() {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"session is expected to be updated but is not",
		)
	}
	return nil
}
