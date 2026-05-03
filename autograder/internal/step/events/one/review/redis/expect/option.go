package expect

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/review/count"
)

// Option represents the functional option
// to configure the Exectations instance on its creation.
type Option func(e *Expectations)

// WithCounts sets the Counts to the Expectations instance on creation.
func WithCounts(c count.Counts) Option {
	return func(e *Expectations) {
		e.counts = c
	}
}

// WithTTL sets the time-to-live of the value
// to the Expectations instance on creation.
func WithTTL(ttl time.Duration) Option {
	return func(e *Expectations) {
		e.ttl = ttl
	}
}
