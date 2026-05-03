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
		Reviews []json.RawMessage `json:"reviews"`
		Count   int               `json:"count"`
	}
	jsonbody.NewBody(body).MustParseIn(&v)
	var reviews []event.Review
	if len(v.Reviews) > 0 {
		reviews = make([]event.Review, 0, len(v.Reviews))
		for _, jsn := range v.Reviews {
			reviews = append(reviews, event.MustParseJSON(jsn))
		}
	}
	return NewBody(
		reviews,
		v.Count,
	)
}
