package endpoint_test

import (
	"time"

	cookie "github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/reviews/events/create/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/rating"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
	cookieassert "github.com/sitnikovik/ndbx/autograder/internal/expect/http/cookie"
	cookiexpct "github.com/sitnikovik/ndbx/autograder/internal/expect/http/cookie/expectation"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/http/response/expectation"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	eventfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/event"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
	sidfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/cookie/session/id"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
	usrsession "github.com/sitnikovik/ndbx/autograder/internal/user/session"
)

var (
	// descFixture is a fixture used in the tests cases.
	descFixture = step.NewDesc(
		"Test step",
		"Test description",
	)
	// eventFixture is a fixture used in the tests cases.
	eventFixture = eventfx.NewBirthdayParty(
		event.NewDates(
			timex.MustRFC3339("2026-03-31T15:00:00Z"),
			timex.MustRFC3339("2026-03-31T23:00:00Z"),
		),
		timex.MustRFC3339("2026-03-14T12:31:00Z"),
		userfx.NewSamwiseGamgee(),
	)
	bodyFixture = body.NewBody(
		body.WithComment("test comment"),
		body.WithRating(rating.Five),
	)
	// varsFixture is the step variables used in the tests cases.
	varsFixture = func() step.Variables {
		vv := step.NewVariables()
		vv.Set(
			cookie.Name,
			sidfx.OK,
		)
		vv.Set(
			cookie.Name,
			sidfx.OK,
		)
		vv.Set(
			variable.SessionTTL,
			3600*time.Second,
		)
		vv.Set(
			eventFixture.Hash(),
			"13298",
		)
		return vv
	}()
	xpctFixture = expectation.NewExpectations(
		expectation.WithAsserts(
			response.AssertCreatedStatus,
			response.AssertNotEmptyContent,
		),
		expectation.WithCookies(
			cookiexpct.NewExpectations(
				cookie.Name,
				cookiexpct.WithAsserts(
					cookieassert.AssertExistsMaxAge,
					cookieassert.AssertExistsHTTPOnly,
				),
				cookiexpct.WithAssertsValueFn(
					usrsession.Validate,
				),
			),
		),
	)
)
