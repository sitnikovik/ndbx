package body

import (
	"fmt"
	"io"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

// MustParseBody reads the HTTP response body represents the target user.
//
// Panics if parsing the body fails or if the required fields are missing.
func MustParseBody(body io.ReadCloser) Body {
	bb, err := io.ReadAll(body)
	if err != nil {
		panic(fmt.Sprintf("failed to read body to be parsed to user: %v", err))
	}
	return NewBody(
		user.MustParseJSON(bb),
	)
}
