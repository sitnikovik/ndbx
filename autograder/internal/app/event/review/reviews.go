package review

import common "github.com/sitnikovik/ndbx/autograder/internal/app/review/count"

// Reviews represents the reviews left by users for an event.
type Reviews struct {
	// counts is the counts of the reviews left by users for an event.
	counts common.Counts
}

// NewReviews creates a new Reviews instance
// that can be configured with the provided options.
func NewReviews(opts ...Option) Reviews {
	r := Reviews{}
	for _, opt := range opts {
		opt(&r)
	}
	return r
}

// With copies the instance and returns a new one with provided options.
func (r Reviews) With(opts ...Option) Reviews {
	cop := r
	for _, opt := range opts {
		opt(&cop)
	}
	return cop
}

// Counts returns the counts of the reviews left by users for an event.
func (r Reviews) Counts() common.Counts {
	return r.counts
}
