package event

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

// NewTestEvent creates a fixture of the Event used in tests.
func NewTestEvent() event.Event {
	return event.NewEvent(
		event.NewID("1"),
		event.NewContent("Test event", "Description"),
		event.NewLocation("NY"),
		event.NewCreated(
			timex.MustRFC3339("2025-01-03T11:00:00Z"),
			user.NewIdentity(user.ID("123")),
		),
		event.NewDates(
			timex.MustRFC3339("2025-01-11T11:00:00Z"),
			timex.MustRFC3339("2025-01-11T23:00:00Z"),
		),
	)
}
