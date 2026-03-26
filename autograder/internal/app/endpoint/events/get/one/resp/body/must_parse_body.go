package body

import (
	"fmt"
	"io"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
)

// MustParseBody reads the HTTP response body represents the target event.
//
// Panics if parsing the body fails or if the required fields are missing.
func MustParseBody(body io.ReadCloser) Body {
	bb, err := io.ReadAll(body)
	if err != nil {
		panic(fmt.Sprintf("failed to read body to be parsed to event: %v", err))
	}
	return NewBody(
		event.MustParseJSON(bb),
	)
}
