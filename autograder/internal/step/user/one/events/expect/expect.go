package expect

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/reaction"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/review"
)

// Expectations holds the expectations we need to check in the Step.
type Expectations struct {
	// reactions is the event reactions for each event to expect.
	reactions []reaction.Reactions
	// reviews is the event reviews for each event to expect.
	reviews []review.Reviews
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

// HasReviews defines are reviews set in the Expectations instance.
func (e Expectations) HasReviews() bool {
	return e.reviews != nil
}

// Reviews returns event Reviews to expect.
func (e Expectations) Reviews() []review.Reviews {
	return e.reviews
}
