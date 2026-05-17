package expect

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
)

// Expectations holds the expectations we need to check in the Step.
type Expectations struct {
	// events is the list of events to expect.
	events []event.Event
	// ttl is the duration of the list to expect.
	ttl time.Duration
}

// NewExpectations creates a new Expectations instance with the given options.
func NewExpectations(opt Option, opts ...Option) Expectations {
	e := Expectations{}
	opt(&e)
	for _, o := range opts {
		o(&e)
	}
	return e
}

// HasEvents defines if events are set in the Expectations.
func (e Expectations) HasEvents() bool {
	return e.events != nil
}

// Events returns the list of events to expect.
func (e Expectations) Events() []event.Event {
	return e.events
}

// TTL returns the duration of the list to expect.
func (e Expectations) TTL() time.Duration {
	return e.ttl
}
