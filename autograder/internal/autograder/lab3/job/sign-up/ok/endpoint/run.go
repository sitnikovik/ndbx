package endpoint

import (
	"bytes"
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	request "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/users/post/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes the sign-up endpoint test step,
// sending a POST request to the sign-up endpoint and validating the response.
//
// Sets values of name, username, and password used in the sign-up process
// as step variables to use them in the next steps.
func (s *Step) Run(
	_ context.Context,
	vars step.Variables,
) error {
	rsp, err := s.cli.PostJSON(
		endpoint.
			NewEndpoint(s.baseURL).
			SignUp(),
		bytes.NewBuffer(
			request.
				NewBody(
					lab3.NewTestUser(),
					lab3.TestUserPassword,
				).
				MustBytes(),
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
		response.AssertCreatedStatus,
		response.AssertEmptyContent,
	)
	if err != nil {
		return errs.Wrap(
			err,
			"got unexpected response",
		)
	}
	sess := session.MustParseSession(rsp.Cookies())
	err = sess.Validate()
	if err != nil {
		return errs.Wrap(
			err,
			"got invalid session cookie",
		)
	}
	vars.Set(
		variable.User,
		lab3.NewTestUser(),
	)
	vars.Set(
		variable.UserPassword,
		lab3.TestUserPassword,
	)
	vars.Set(
		session.Name,
		sess.String(),
	)
	return nil
}
