package review

import common "github.com/sitnikovik/ndbx/autograder/internal/app/review/count"

// Option represents a functional option
// to configure the Reviews instance on its creation.
type Option func(*Reviews)

// WithCounts sets the Counts to the Reviews instance.
func WithCounts(counts common.Counts) Option {
	return func(r *Reviews) {
		r.counts = counts
	}
}
