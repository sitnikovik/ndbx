package body

import "github.com/sitnikovik/ndbx/autograder/internal/app/rating"

// Option represents a functional option
// to configure the Body instance on its creation.
type Option func(*Body)

// WithRating sets the rating for the Body instance on creation.
func WithRating(rate rating.Rating) Option {
	return func(b *Body) {
		b.rating = rate
	}
}

// WithComment sets the comment for the Body instance on creation.
func WithComment(comment string) Option {
	return func(b *Body) {
		b.comment = comment
	}
}
