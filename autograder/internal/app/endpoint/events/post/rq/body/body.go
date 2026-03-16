package body

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

// Body represents the request body for the authentication endpoint.
type Body struct {
	e event.Event
}

// NewBody creates a new Body instance with the provided event data.
func NewBody(event event.Event) Body {
	return Body{
		e: event,
	}
}

// MustBytes returns the JSON representation of the Body as a byte slice.
//
// Panics if marshalling the data fails.
func (b Body) MustBytes() []byte {
	m := make(map[string]any, 10)
	m["title"] = b.e.Content().Title()
	m["address"] = b.e.Location().Address()
	m["started_at"] = b.e.Dates().StartedAt().Format(time.RFC3339)
	m["finished_at"] = b.e.Dates().FinishedAt().Format(time.RFC3339)
	if v := b.e.Content().Description(); v != "" {
		m["description"] = v
	}
	if v := b.e.Quantity().Max(); v > 0 {
		m["max_attendees"] = v
	}
	if v := b.e.Quantity().Min(); v > 0 {
		m["min_attendees"] = v
	}
	bb, err := json.Marshal(m)
	if err != nil {
		panic(errors.Join(errs.ErrMarshallFailed, err))
	}
	return bb
}
