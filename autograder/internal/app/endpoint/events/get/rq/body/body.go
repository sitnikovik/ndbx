package body

import (
	"net/url"

	"github.com/sitnikovik/fluxhttp/query"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/include"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/pagination"
	"github.com/sitnikovik/ndbx/autograder/internal/console"
)

// Body represents the request body for the list events endpoint.
type Body struct {
	// created holds the creation parameters for filtering events.
	created Created
	// content holds the specific query parameters for filtering events.
	content Content
	// loc holds the location parameters for filtering events.
	loc Location
	// costs holds the cost parameters for filtering events.
	costs Costs
	// inc holds the include parameters for the request.
	inc include.Include
	// pg holds the pagination parameters for the request.
	pg pagination.Pagination
}

// NewBody creates a new Body instance with the provided options.
func NewBody(opts ...Option) Body {
	b := Body{}
	for _, opt := range opts {
		opt(&b)
	}
	return b
}

// URLQuery converts the Body into url.Values by merging the URL queries of its fields.
func (b Body) URLQuery() url.Values {
	q := make(url.Values)
	query.MergeInto(
		q,
		b.created.URLQuery(),
		b.content.URLQuery(),
		b.loc.URLQuery(),
		b.costs.URLQuery(),
		b.inc.URLQuery(),
		b.pg.URLQuery(),
	)
	console.Log("query: %s", q.Encode())
	return q
}
