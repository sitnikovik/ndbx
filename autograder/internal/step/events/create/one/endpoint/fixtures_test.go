package endpoint_test

import (
	session "github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	cookieasserts "github.com/sitnikovik/ndbx/autograder/internal/expect/http/cookie"
	cookiexpct "github.com/sitnikovik/ndbx/autograder/internal/expect/http/cookie/expectation"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response/expectation"
	sessions "github.com/sitnikovik/ndbx/autograder/internal/user/session"
)

// NewExpectationsFx creates a new instance of HTTP response expecations.
func NewExpectationsFx() expectation.Expectations {
	return expectation.NewExpectations(
		expectation.WithAsserts(
			response.AssertCreatedStatus,
			response.AssertNotEmptyContent,
		),
		expectation.WithCookies(
			cookiexpct.NewExpectations(
				session.Name,
				cookiexpct.WithAssertsValueFn(
					sessions.Validate,
				),
				cookiexpct.WithAsserts(
					cookieasserts.AssertExistsMaxAge,
					cookieasserts.AssertExistsHTTPOnly,
				),
			),
		),
	)
}
