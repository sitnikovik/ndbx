package expect

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
)

// Option represents the functional option
// to configure the Exectations instance on its creation.
type Option func(*Expectations)

// WithTTL sets the time-to-live of the value
// to the Expectations instance on creation.
func WithTTL(ttl time.Duration) Option {
	return func(e *Expectations) {
		e.ttl = ttl
	}
}

// WithEvents sets the events to the Expectations instance on creation.
func WithEvents(ee ...event.Event) Option {
	return func(e *Expectations) {
		e.events = ee
	}
}

// WithNoEvents sets that there is no any event in response to expect.
func WithNoEvents() Option {
	return func(e *Expectations) {
		e.events = []event.Event{}
	}
}
