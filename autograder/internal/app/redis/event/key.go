package event

import "github.com/sitnikovik/ndbx/autograder/internal/app/event"

// Key returns the key for the event set in Redis.
func Key(id event.ID) string {
	return "event:" + id.String()
}
