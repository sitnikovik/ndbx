package body

import (
	"encoding/json"
	"io"

	"github.com/sitnikovik/ndbx/autograder/internal/app/review/event"
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
	var reviews []event.Review
	if len(v.Events) > 0 {
		reviews = make([]event.Review, 0, len(v.Events))
		for _, jsn := range v.Events {
			reviews = append(reviews, event.MustParseJSON(jsn))
		}
	}
	return NewBody(
		reviews,
		v.Count,
	)
}
