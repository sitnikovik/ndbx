package body

import (
	"net/url"

	"github.com/sitnikovik/fluxhttp/query"

	rangeof "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/range-of"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

// Created represents the request body to filter events by creation data.
type Created struct {
	// at represents the date range to get the events for.
	at rangeof.Dates
	// by represents the user to get the events for.
	by user.Identity
}

// URLQuery converts the Created into url.Values.
func (c Created) URLQuery() url.Values {
	q := make(url.Values, 2)
	if v := c.by.ID(); !v.Empty() {
		q.Set("user_id", v.String())
	}
	if v := c.by.Username(); v != "" {
		q.Set("user", v)
	}
	query.MergeInto(q, c.at.URLQuery())
	return q
}
