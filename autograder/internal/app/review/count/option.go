package count

import "github.com/sitnikovik/ndbx/autograder/internal/app/rating"

// Option represents a functional option for configuring the Counts instance.
type Option func(c *Counts)

// WithRating sets the rating for the Counts instance on creation.
func WithRating(r rating.Rating) Option {
	return func(c *Counts) {
		c.rating = r
	}
}

// WithCount sets the count for the Counts instance on creation.
func WithCount(n int) Option {
	return func(c *Counts) {
		c.count = n
	}
}
