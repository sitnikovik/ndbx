package reaction

import "github.com/sitnikovik/ndbx/autograder/internal/app/cassandra/event/reaction/filter"

// Option represents a functional option
// to configure filtering the Reactions instance.
type Option func(*Reactions)

// WithLimit sets the limit to restrict
// the number of Reactions got from database.
func WithLimit(n int) Option {
	return func(r *Reactions) {
		r.limit = n
	}
}

// WithFilter sets the filter for reaction selecting.
func WithFilter(f filter.Filter) Option {
	return func(r *Reactions) {
		r.ftr = f
	}
}
