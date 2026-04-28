package event

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	"github.com/sitnikovik/ndbx/autograder/internal/app/rating"
)

// Review represents the event review that users have given.
type Review struct {
	// stamp holds creation metadata.
	stamp creation.Stamp
	// updatedAt holds the last update time.
	updatedAt time.Time
	// comment is the review comment.
	comment string
	// event is metadata describes what the event has been reviewed.
	event Event
	// id is the review idendtifier.
	id string
	// rating is the review rating that users have given.
	rating rating.Rating
}

// NewReview creates a new Review instance.
func NewReview(
	id string,
	stamp creation.Stamp,
	ev Event,
	comment string,
	rating rating.Rating,
	opts ...Option,
) Review {
	r := Review{
		id:      id,
		stamp:   stamp,
		event:   ev,
		comment: comment,
		rating:  rating,
	}
	for _, opt := range opts {
		opt(&r)
	}
	return r
}

// Created returns creation metadata.
func (r Review) Created() creation.Created {
	return r.stamp.Created()
}

// UpdatedAt returns the last update time.
func (r Review) UpdatedAt() time.Time {
	return r.updatedAt
}

// Comment returns the review comment.
func (r Review) Comment() string {
	return r.comment
}

// Event returns metadata describes what the event has been reviewed.
func (r Review) Event() Event {
	return r.event
}

// ID returns the review idendtifier.
func (r Review) ID() string {
	return r.id
}

// Rating returns the review rating that users have given.
func (r Review) Rating() rating.Rating {
	return r.rating
}
