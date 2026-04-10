package expect

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/reaction"
)

// Option represents the functional option
// to configure the Exectations instance on its creation.
type Option func(e *Expectations)

// WithEvents sets the Events to the Expectations instance on creation.
func WithEvents(ee ...event.Event) Option {
	return func(e *Expectations) {
		e.events = ee
	}
}

// WithNoEvents sets that there is no any events in response to expect.
func WithNoEvents() Option {
	return func(e *Expectations) {
		e.events = []event.Event{}
	}
}

// WithReactions sets event reactions to the Expectations instance on creation.
//
// The list is supposed to match the length of events to expect.
func WithReactions(rr ...reaction.Reactions) Option {
	return func(e *Expectations) {
		e.reactions = rr
	}
}

// WithNoReactions sets that there is no any event reactions in response to expect.
func WithNoReactions() Option {
	return func(e *Expectations) {
		e.reactions = []reaction.Reactions{}
	}
}
