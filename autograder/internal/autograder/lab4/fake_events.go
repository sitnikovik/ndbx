package lab4

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

// FakeEvents generates and returns a slice of fake events for autograder purposes.
//
// The number of events generated is determined by the `limit` parameter,
// and each event is associated with a user ID from the `userIDs` slice.
//
// It shuffles the generated events to ensure randomness in their order by user IDs.
func FakeEvents(
	limit int,
	userIDs []user.ID,
) []event.Event {
	res := make([]event.Event, len(userIDs)*limit)
	usersLen := len(userIDs)
	for i := range limit {
		num := fmt.Sprint(i + 1)
		createdAt := timex.
			MustRFC3339("2025-02-01T09:00:00Z").
			Add(
				time.Duration(rand.IntN(10))*time.Hour +
					time.Duration(rand.IntN(60))*time.Minute,
			)
		res[i] = event.NewEvent(
			event.NewID(""),
			event.NewContent(
				"Title "+num,
				"Description for event "+num,
			),
			event.NewLocation(""),
			event.NewCreated(
				createdAt,
				user.NewIdentity(userIDs[rand.IntN(usersLen-1)]),
			),
			event.NewDates(
				createdAt.Add(2*time.Hour),
				createdAt.Add(4*time.Hour),
			),
		)
	}
	return res
}
