package body

import (
	"net/url"

	"github.com/sitnikovik/fluxhttp/query"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/pagination"
)

// Body represents the request body for the list events endpoint.
type Body struct {
	// content holds the specific query parameters for filtering events.
	content Content
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
	query.MergeInto(q, b.content.URLQuery(), b.pg.URLQuery())
	return q
}
