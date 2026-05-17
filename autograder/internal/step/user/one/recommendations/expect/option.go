package expect

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	resp "github.com/sitnikovik/ndbx/autograder/internal/expect/http/response/expectation"
)

// Option represents the functional option
// to configure the Exectations instance on its creation.
type Option func(e *Expectations)

// WithReviews sets the list of events
// to the Expectations instance on its creation.
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

// WithResponse sets the response
// to the Expectations instance on its creation.
func WithResponse(r resp.Expectations) Option {
	return func(e *Expectations) {
		e.resp = r
	}
}
