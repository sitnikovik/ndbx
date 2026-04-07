package endpoint

import (
	"github.com/sitnikovik/ndbx/autograder/internal/step/user/one/events/expect"
)

// Option represents the functional option
// to configure the Step instance on its creation.
type Option func(s *Step)

// WithExpectations sets the Expectations
// to the Step instance on creation.
func WithExpectations(exp expect.Expectations) Option {
	return func(s *Step) {
		s.expect = exp
	}
}
