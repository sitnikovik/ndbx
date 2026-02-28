package endpoint

import (
	"context"
	"errors"
	"fmt"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	consts "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/cookie"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	numbersExpected "github.com/sitnikovik/ndbx/autograder/internal/expect/numbers"
	"github.com/sitnikovik/ndbx/autograder/internal/http/response/cookie"
	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	"github.com/sitnikovik/ndbx/autograder/internal/user/session"
)

// Run performs the HTTP session step by sending a POST request to the specified URL,
// validating the response, and extracting the session cookie.
func (s *Step) Run(
	_ context.Context,
	vars step.Variables,
) error {
	url := endpoint.
		NewEndpoint(s.url).
		Session()
	resp, err := s.cli.PostJSON(url, nil)
	if err != nil {
		return errors.Join(errs.ErrHTTPFailed, err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			panic(fmt.Errorf("failed to close response body: %w", err))
		}
	}()
	err = numbersExpected.AssertEquals(201, resp.StatusCode)
	if err != nil {
		return errors.Join(err, errs.ErrInvalidHTTPStatus)
	}
	err = numbersExpected.AssertEmpty(resp.ContentLength)
	if err != nil {
		return errs.Wrap(err, "response content length")
	}
	ckname := consts.SessionName
	ck := cookie.
		NewCookies(resp.Cookies()).
		MustGet(ckname)
	sid := ck.Value
	if sid == "" {
		return errs.Wrap(
			errs.ErrMissedCookie,
			"expect %s cookie to have a value",
			log.String(ckname),
		)
	}
	if !ck.HttpOnly {
		return errs.Wrap(
			errs.ErrMissedCookie,
			"expect %s cookie to have http only flag",
			log.String(ckname),
		)
	}
	if ck.MaxAge <= 0 {
		return errs.Wrap(
			errs.ErrMissedCookie,
			"expect %s cookie to have MaxAge flag",
			log.String(ckname),
		)
	}
	err = session.Validate(sid)
	if err != nil {
		return errors.Join(errs.ErrExpectationFailed, err)
	}
	vars.Set(ckname, sid)
	return nil
}
