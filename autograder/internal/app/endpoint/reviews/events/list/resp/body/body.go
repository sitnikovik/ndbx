package body

import "github.com/sitnikovik/ndbx/autograder/internal/app/review/event"

// Body represents the HTTP response body of the list of reviews endpoint.
type Body struct {
	// reviews is the list of the reviews got from response.
	reviews []event.Review
	// count is the total number of reviews.
	count int
}

// NewBody creates a new Body instance.
func NewBody(reviews []event.Review, count int) Body {
	return Body{
		reviews: reviews,
		count:   count,
	}
}

// Reviews returns a list of the reviews got from response.
func (b Body) Reviews() []event.Review {
	return b.reviews
}

// Count returns total count of the list.
func (b Body) Count() int {
	return b.count
}
