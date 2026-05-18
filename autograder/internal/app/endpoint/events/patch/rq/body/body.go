package body

import (
	"encoding/json"
	"net/url"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// Body represents the request body to updated an event.
type Body struct {
	// tags is the list of tags to update
	tags []string
	// category is the category of the event to be set.
	category string
	// city is the city of the event location to be set.
	city string
	// price is the price of the event to be set.
	price uint
	// cascade defines whether the patch is to be cascade update.
	cascade bool
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
	m := make(map[string]any, 4)
	if b.tags != nil {
		m["tags"] = b.tags
	}
	if b.category != "" {
		m["category"] = b.category
	}
	if b.city != "" {
		m["city"] = b.city
	}
	if b.price != 0 {
		m["price"] = b.price
	}
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

// URLQuery converts the Body into url.Values.
func (b Body) URLQuery() url.Values {
	q := make(url.Values, 1)
	if b.cascade {
		q.Set("cascade", "true")
	}
	return q
}
