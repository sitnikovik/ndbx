package expect

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	resp "github.com/sitnikovik/ndbx/autograder/internal/expect/http/response/expectation"
)

// Expectations holds the expectations we need to check in the Step.
type Expectations struct {
	// events is the list of events to expect.
	events []event.Event
	// resp is the response to expect.
	resp resp.Expectations
}

// NewExpectations creates a new Expectations instance.
func NewExpectations(opt Option, opts ...Option) Expectations {
	e := Expectations{}
	opt(&e)
	for _, o := range opts {
		o(&e)
	}
	return e
}

// HasEvents defines are events set in the Expectations instance.
func (e Expectations) HasEvents() bool {
	return e.events != nil
}

// Events returns events to expect
func (e Expectations) Events() []event.Event {
	return e.events
}

// Response returns the response to expect.
func (e Expectations) Response() resp.Expectations {
	return e.resp
}
