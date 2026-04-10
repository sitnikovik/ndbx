package reaction

import common "github.com/sitnikovik/ndbx/autograder/internal/app/reaction/count"

// Reactions represents the reactions left by users for an event.
type Reactions struct {
	// counts is the counts of the reactions left by users for an event.
	counts common.Counts
}

// NewReactions creates a new Reactions instance
// that can be configured with the provided options.
func NewReactions(opts ...Option) Reactions {
	r := Reactions{}
	for _, opt := range opts {
		opt(&r)
	}
	return r
}

// With copies the instance and returns a new one with provided options.
func (r Reactions) With(opts ...Option) Reactions {
	cop := r
	for _, opt := range opts {
		opt(&cop)
	}
	return cop
}

// Counts returns the counts of the reactions left by users for an event.
func (r Reactions) Counts() common.Counts {
	return r.counts
}
