package body

import (
	"encoding/json"

	"github.com/sitnikovik/ndbx/autograder/internal/app/rating"
)

// Body represents the request body for event reviews endpoint.
type Body struct {
	// comment is the comment of the event review.
	comment string
	// rating is the rating of the event review.
	rating rating.Rating
}

// NewBody creates a new Body instance with the given options.
func NewBody(
	opt Option,
	opts ...Option,
) Body {
	r := Body{}
	opt(&r)
	for _, o := range opts {
		o(&r)
	}
	return r
}

// MustBytes returns the JSON representation of the Body as a byte slice.
//
// Panics if marshalling the data fails.
func (b Body) MustBytes() []byte {
	m := make(map[string]any, 2)
	if b.comment != "" {
		m["comment"] = b.comment
	}
	if rate := b.rating; !rate.Empty() {
		m["rating"] = rate.Int()
	}
	bb, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return bb
}
