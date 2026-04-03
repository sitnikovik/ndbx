package event

import "github.com/sitnikovik/ndbx/autograder/internal/app/event"

// Event represents metadata describes what the event has been liked.
type Event struct {
	// id is the of the event has been liked.
	id event.ID
}

// NewEvent creates a new Event instance holds event metadata.
func NewEvent(id event.ID) Event {
	return Event{
		id: id,
	}
}

// ID returns the event id.
func (e Event) ID() event.ID {
	return e.id
}
