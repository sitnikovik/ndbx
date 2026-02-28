package endpoint

import (
	"context"
	"errors"
	"fmt"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	consts "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/cookie"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	numbersExpected "github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/strings"
	"github.com/sitnikovik/ndbx/autograder/internal/http/response/cookie"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	"github.com/sitnikovik/ndbx/autograder/internal/user/session"
)

// Run performs the HTTP health check step by sending a GET request to the specified URL,
// validating the response, and ensuring the session cookie is correctly set.
func (s *Step) Run(
	_ context.Context,
	vars step.Variables,
) error {
	url := endpoint.
		NewEndpoint(s.baseURL).
		Health()
	resp, err := s.cli.Get(url)
	if err != nil {
		return errors.Join(
			errs.ErrHTTPFailed,
			err,
		)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			panic(fmt.Errorf("failed to close response body: %w", err))
		}
	}()
	err = numbersExpected.AssertEquals(200, resp.StatusCode)
	if err != nil {
		return errors.Join(
			errs.ErrInvalidHTTPStatus,
			errs.ErrExpectationFailed,
			err,
		)
	}
	ckname := consts.SessionName
	ck := cookie.
		NewCookies(resp.Cookies()).
		MustGet(ckname)
	sid := ck.Value
	err = session.Validate(sid)
	if err != nil {
		return errs.Wrap(
			errors.Join(
				errs.ErrExpectationFailed,
				err,
			),
			"invalid session id in cookie",
		)
	}
	err = strings.AssertEquals(
		vars.
			MustGet(ckname).
			AsString(),
		sid,
	)
	if err != nil {
		return errs.Wrap(
			errors.Join(
				errs.ErrExpectationFailed,
				err,
			),
			"session id in cookie does not match expected value",
		)
	}
	if !ck.HttpOnly {
		return errs.Wrap(
			errs.ErrExpectationFailed,
			"expect %s cookie to have http only flag",
			log.String(ckname),
		)
	}
	return nil
}
