package event

import "github.com/sitnikovik/ndbx/autograder/internal/app/event"

// Event represents an event in the Neo4j database
// of the target application.
type Event struct {
	// title is the title of the event.
	title string
	// id is the unique identifier of the event.
	id event.ID
}

// NewEvent creates a new Event stored in Neo4j
// with the given event ID and title.
func NewEvent(id event.ID, title string) Event {
	return Event{
		id:    id,
		title: title,
	}
}

// ID returns the unique identifier of the event.
func (e Event) ID() event.ID {
	return e.id
}

// Title returns the title of the event.
func (e Event) Title() string {
	return e.title
}
