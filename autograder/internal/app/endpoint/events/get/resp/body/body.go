package body

import "github.com/sitnikovik/ndbx/autograder/internal/app/event"

// Body represents the HTTP response body of the list of events endpoint.
type Body struct {
	// events is the list of the events got from response.
	events []event.Event
	// n is total count of the list.
	n int
}

// NewBody creates a new Body instance.
func NewBody(events []event.Event, n int) Body {
	return Body{
		events: events,
		n:      n,
	}
}

// Events returns a list of the events got from response.
func (b Body) Events() []event.Event {
	return b.events
}

// Count returns total count of the list.
func (b Body) Count() int {
	return b.n
}
