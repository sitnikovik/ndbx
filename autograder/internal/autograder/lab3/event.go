package lab3

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

// NewTestEvent creates and returns a new event to be used in the autograder.
func NewTestEvent() event.Event {
	return event.NewEvent(
		event.NewID("000000000000000000000001"),
		event.NewContent("Title", "Description"),
		event.NewLocation("City, Country, Street, 123"),
		event.NewCreated(
			timex.MustParse(time.RFC3339, "2025-02-01T09:00:00Z"),
			user.NewIdentity(
				user.NewID("000000000000000000000123"),
			),
		),
		event.NewDates(
			timex.MustParse(time.RFC3339, "2025-02-01T11:00:00Z"),
			timex.MustParse(time.RFC3339, "2025-02-01T13:00:00Z"),
		),
	)
}
