package expect

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/reaction"
)

// Option represents the functional option
// to configure the Exectations instance on its creation.
type Option func(e *Expectations)

// WithReactions sets event reactions to the Expectations instance on creation.
//
// The list is supposed to match the length of events to expect.
func WithReactions(rr []reaction.Reactions) Option {
	return func(e *Expectations) {
		e.reactions = rr
	}
}
