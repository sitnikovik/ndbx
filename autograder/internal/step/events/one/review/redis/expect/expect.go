package expect

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/review/count"
)

// Expectations holds the expectations we need to check in the Step.
type Expectations struct {
	// counts is the counts of the reviews left by users for an event.
	counts count.Counts
	// ttl is the time to live for the value.
	ttl time.Duration
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

// Counts returns counts to expect.
func (e Expectations) Counts() count.Counts {
	return e.counts
}

// TTL returns time-to-live of the value to expect.
func (e Expectations) TTL() time.Duration {
	return e.ttl
}
