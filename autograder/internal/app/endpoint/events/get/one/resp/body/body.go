package body

import "github.com/sitnikovik/ndbx/autograder/internal/app/event"

// Body represents the HTTP response body of the event endpoint.
type Body struct {
	// event is the event got by the endpoint.
	event event.Event
}

// NewBody creates a new Body instance.
func NewBody(ev event.Event) Body {
	return Body{
		event: ev,
	}
}

// Event returns the Event got by the endpoint.
func (b Body) Event() event.Event {
	return b.event
}
