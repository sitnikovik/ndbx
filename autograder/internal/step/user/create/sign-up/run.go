package endpoint

import (
	"bytes"
	"context"
	"errors"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	request "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/users/post/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/user/collection"
	dockey "github.com/sitnikovik/ndbx/autograder/internal/app/mongo/user/doc/key"
	mongodoc "github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

// Run executes the sign-up endpoint test step,
// sending a POST request to the sign-up endpoint and validating the response.
//
// Sets values of name, username, and password used in the sign-up process
// as step variables to use them in the next steps.
func (s *Step) Run(
	ctx context.Context,
	vars step.Variables,
) error {
	rsp, err := s.cli.PostJSON(
		endpoint.
			NewEndpoint(s.baseURL).
			SignUp(),
		bytes.NewBuffer(
			request.
				NewBody(
					s.user,
					s.pwd,
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
	users, err := s.mongo.AllBy(
		ctx,
		collection.Name,
		mongodoc.NewKVs(
			mongodoc.NewKV(
				dockey.Username,
				s.user.Username(),
			),
		),
	)
	if err != nil {
		return errs.WrapJoin(
			"failed to get users from MongoDB",
			errs.ErrExternalDependencyFailed,
			err,
		)
	}
	err = numbers.AssertEquals(1, users.Len())
	if err != nil {
		return err
	}
	vars.Set(
		session.Name,
		sess.String(),
	)
	vars.Set(
		s.user.Hash(),
		users.First().ID(),
	)
	return nil
}
