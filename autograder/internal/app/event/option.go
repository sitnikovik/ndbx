package event

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/reaction"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/review"
	"github.com/sitnikovik/ndbx/autograder/internal/app/rating"
	"github.com/sitnikovik/ndbx/autograder/internal/app/reaction/count"
	reviewCount "github.com/sitnikovik/ndbx/autograder/internal/app/review/count"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

// Option represents a functional option for configuring an Event.
type Option func(*Event)

// WithID sets the unique identifier for the event.
func WithID(id ID) Option {
	return func(e *Event) {
		e.id = id
	}
}

// WithCreatedBy sets the user's identification for the event.
func WithCreatedBy(id user.Identity) Option {
	return func(e *Event) {
		e.created.by = id
	}
}

// WithQuantity sets the quantity of attendees for the event in the Event.
func WithQuantity(q Quantity) Option {
	return func(b *Event) {
		b.qty = q
	}
}

// WithCosts sets the cost information related to the event.
func WithCosts(costs Costs) Option {
	return func(e *Event) {
		e.costs = costs
	}
}

// WithReactions set the reactions to the event.
func WithReactions(reactions reaction.Reactions) Option {
	return func(e *Event) {
		e.reactions = reactions
	}
}

// WithLikes set the counters of the like reactions for the event.
func WithLikes(n uint64) Option {
	return func(e *Event) {
		e.reactions = e.reactions.With(
			reaction.WithCounts(
				e.reactions.
					Counts().
					With(
						count.WithLikes(n),
					),
			),
		)
	}
}

// WithDislikes set the counters of the dislike reactions for the event.
func WithDislikes(n uint64) Option {
	return func(e *Event) {
		e.reactions = e.reactions.With(
			reaction.WithCounts(
				e.reactions.
					Counts().
					With(
						count.WithDislikes(n),
					),
			),
		)
	}
}

// WithReviews set the reviews to the event.
func WithReviews(reviews review.Reviews) Option {
	return func(e *Event) {
		e.reviews = reviews
	}
}

// WithRating sets the rating for the event.
func WithRating(r rating.Rating) Option {
	return func(e *Event) {
		e.reviews = e.reviews.With(
			review.WithCounts(
				e.reviews.
					Counts().
					With(
						reviewCount.WithRating(r),
					),
			),
		)
	}
}
