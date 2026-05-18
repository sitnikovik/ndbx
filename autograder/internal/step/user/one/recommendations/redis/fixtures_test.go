package redis_test

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	"github.com/sitnikovik/ndbx/autograder/internal/step/user/one/recommendations/redis/expect"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

// NewDescFx returns a new step desc fixture.
func NewDescFx() step.Desc {
	return step.NewDesc("Title", "Description")
}

// NewExpectationsFx returns a new expectations fixture.
func NewExpectationsFx() expect.Expectations {
	return expect.NewExpectations(
		expect.WithEvents(
			event.NewEvent(
				event.NewID("1"),
				event.NewContent(
					"test title",
					"test description",
				),
				event.NewLocation("test location"),
				event.NewCreated(
					timex.MustRFC3339("2024-01-01T00:00:00Z"),
					user.NewIdentity("test_user"),
				),
				event.NewDates(
					timex.MustRFC3339("2024-01-01T01:00:00Z"),
					timex.MustRFC3339("2024-01-01T02:00:00Z"),
				),
			),
		),
		expect.WithTTL(1*time.Minute),
	)
}

