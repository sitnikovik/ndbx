package event

// Event represents an event.
type Event struct {
	// created contains the creation details of the event.
	created Created
	// dates contains the start and finish dates of the event.
	dates Dates
	// content contains the content info of the event.
	content Content
	// loc is the location of the event.
	loc Location
	// qty represents the quantity of attendees for the event.
	qty Quantity
	// id is the unique identifier for the event.
	id ID
}

// NewEvent creates a new Event instance.
func NewEvent(
	id ID,
	cont Content,
	loc Location,
	created Created,
	dates Dates,
	opts ...Option,
) Event {
	e := Event{
		id:      id,
		content: cont,
		created: created,
		loc:     loc,
		dates:   dates,
	}
	for _, opt := range opts {
		opt(&e)
	}
	return e
}

// Content returns the content info of the event.
func (e Event) Content() Content {
	return e.content
}

// Created returns the creation details of the event.
func (e Event) Created() Created {
	return e.created
}

// Location returns the location of the event.
func (e Event) Location() Location {
	return e.loc
}

// Dates returns the start and finish dates of the event.
func (e Event) Dates() Dates {
	return e.dates
}

// Quantity returns the quantity of attendees for the event.
func (e Event) Quantity() Quantity {
	return e.qty
}

// ID returns the unique identifier for the event.
func (e Event) ID() ID {
	return e.id
}
