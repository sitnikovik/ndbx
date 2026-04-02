package event

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event/reaction"
)

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
	// costs contains the cost information related to the event.
	costs Costs
	// qty represents the quantity of attendees for the event.
	qty Quantity
	// reactions represents counters of the users' reactions of the event.
	reactions reaction.Reactions
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

// Costs returns the cost information related to the event.
func (e Event) Costs() Costs {
	return e.costs
}

// Dates returns the start and finish dates of the event.
func (e Event) Dates() Dates {
	return e.dates
}

// Quantity returns the quantity of attendees for the event.
func (e Event) Quantity() Quantity {
	return e.qty
}

// Reactions returns the counters of the reactions for the event.
func (e Event) Reactions() reaction.Reactions {
	return e.reactions
}

// ID returns the unique identifier for the event.
func (e Event) ID() ID {
	return e.id
}

// Hash represents the event as a hash.
func (e Event) Hash() string {
	title := e.Content().Title()
	createdAt := e.Created().At()
	if title == "" && createdAt.IsZero() {
		return ""
	}
	hash := md5.Sum([]byte(title + createdAt.Format(time.RFC3339)))
	return hex.EncodeToString(hash[:])
}
