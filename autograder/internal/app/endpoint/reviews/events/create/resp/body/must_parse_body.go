package body

import (
	"io"

	jsonbody "github.com/sitnikovik/ndbx/autograder/internal/json/body"
)

// MustParseBody reads the HTTP response body and extracts the event ID.
//
// Panics if parsing the body fails or if the required fields are missing.
func MustParseBody(body io.ReadCloser) Body {
	var v struct {
		ID string `json:"id"`
	}
	jsonbody.NewBody(body).MustParseIn(&v)
	return Body{
		id: v.ID,
	}
}
