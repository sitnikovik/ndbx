package event

import "time"

// Option defines a functional option to configure a Review.
type Option func(*Review)

// WithUpdatedAt sets the time when the review was updated.
func WithUpdatedAt(t time.Time) Option {
	return func(r *Review) {
		r.updatedAt = t
	}
}
