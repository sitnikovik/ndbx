package body

import (
	"net/url"

	rangeof "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/range-of"
)

// Costs represents the cost information
// related to the event in the request body.
type Costs struct {
	// entry is the entry price the attendee have to pay.
	entry rangeof.UInts
}

// URLQuery converts the Costs into url.Values.
func (c Costs) URLQuery() url.Values {
	return c.entry.URLQuery()
}
