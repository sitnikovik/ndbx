package count

import "github.com/sitnikovik/ndbx/autograder/internal/app/rating"

// Counts represents counters of reviews for the objects.
type Counts struct {
	// rating is average rating of the object.
	rating rating.Rating
	// count is the number of reviews left.
	count int
}

// NewCounts creates a new Counts instance.
func NewCounts(opts ...Option) Counts {
	c := Counts{}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

// With copies the instance and returns a new one with provided options.
func (c Counts) With(opts ...Option) Counts {
	cop := c
	for _, opt := range opts {
		opt(&cop)
	}
	return cop
}

// Empty defines whether the Counts is empty.
func (c Counts) Empty() bool {
	return c.rating.Empty() && c.count == 0
}

// Rating returns the average rating of the object.
func (c Counts) Rating() rating.Rating {
	return c.rating
}

// Count returns the number of reviews left.
func (c Counts) Count() int {
	return c.count
}
