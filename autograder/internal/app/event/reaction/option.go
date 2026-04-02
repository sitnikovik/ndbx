package reaction

import common "github.com/sitnikovik/ndbx/autograder/internal/app/reaction/count"

// Options represents a functional option
// to configure the Reactions instance on its creation.
type Option func(*Reactions)

// WithCounts sets the Counts to the Reactions instance.
func WithCounts(counts common.Counts) Option {
	return func(r *Reactions) {
		r.counts = counts
	}
}
