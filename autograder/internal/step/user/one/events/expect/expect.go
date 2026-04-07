package expect

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/reaction"
)

// Expectations holds the expectations we need to check in the Step.
type Expectations struct {
	// reactions is the event reactions for each event to expect.
	reactions []reaction.Reactions
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

// HasReactions defines is reactions set in the Expectations instance.
func (e Expectations) HasReactions() bool {
	return e.reactions != nil
}

// Reactions returns event Reactions to expect.
func (e Expectations) Reactions() []reaction.Reactions {
	return e.reactions
}
