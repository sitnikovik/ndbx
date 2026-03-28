package body

import (
	"encoding/json"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// Body represents the request body to updated an event.
type Body struct {
	// category is the category of the event to be set.
	category string
	// city is the city of the event location to be set.
	city string
	// price is the price of the event to be set.
	price uint
}

// NewBody creates a new Body instances with options.
func NewBody(opts ...Option) Body {
	b := Body{}
	for _, opt := range opts {
		opt(&b)
	}
	return b
}

// MustBytes returns the JSON representation of the Body as a byte slice.
//
// Panics if marshalling the data fails.
func (b Body) MustBytes() []byte {
	m := make(map[string]any, 3)
	m["category"] = b.category
	m["city"] = b.city
	m["price"] = b.price
	bb, err := json.Marshal(m)
	if err != nil {
		panic(
			errs.Wrap(
				err,
				"failed to marshal request body for events patch endpoint",
			),
		)
	}
	return bb
}
