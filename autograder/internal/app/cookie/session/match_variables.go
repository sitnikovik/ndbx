package session

import (
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/strings"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// MatchVariables checks if the session matches the expected variables
// and returns an error if any variable does not match.
func (s Session) MatchVariables(vars step.Variables) error {
	err := strings.AssertEquals(
		vars.MustGet(Name).AsString(),
		s.String(),
	)
	if err != nil {
		return errs.Wrap(err, "sid does not match")
	}
	err = numbers.AssertEquals(
		int(vars.MustGet(variable.SessionTTL).AsDuration().Seconds()),
		s.ck.MaxAge,
	)
	if err != nil {
		return errs.Wrap(err, "MaxAge does not match")
	}
	return err
}
