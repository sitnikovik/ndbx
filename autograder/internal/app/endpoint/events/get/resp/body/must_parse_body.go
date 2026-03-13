package body

import (
	"encoding/json"
	"io"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	jsonbody "github.com/sitnikovik/ndbx/autograder/internal/json/body"
)

// MustParseBody reads the HTTP response body represents the list of events.
//
// Panics if parsing the body fails or if the required fields are missing.
func MustParseBody(body io.ReadCloser) Body {
	var v struct {
		Events []json.RawMessage `json:"events"`
		Count  int               `json:"count"`
	}
	jsonbody.NewBody(body).MustParseIn(&v)
	var events []event.Event
	if len(v.Events) > 0 {
		events = make([]event.Event, 0, len(v.Events))
		for _, jsn := range v.Events {
			events = append(events, event.MustParseJSON(jsn))
		}
	}
	return NewBody(
		events,
		v.Count,
	)
}
