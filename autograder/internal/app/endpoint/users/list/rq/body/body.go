package body

import (
	"net/url"

	"github.com/sitnikovik/fluxhttp/query"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/pagination"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

// Body represents the request body for users endpoint.
type Body struct {
	// names holds information how to find the users by thems names.
	names Names
	// id is the information identifies the users.
	id user.Identity
	// pg holds the pagination parameters for the request.
	pg pagination.Pagination
}

// NewBody creates a new Body instance.
func NewBody(opts ...Option) Body {
	b := Body{}
	for _, opt := range opts {
		opt(&b)
	}
	return b
}

// URLQuery converts the Body into url queries
// merging the URL queries of its fields.
func (b Body) URLQuery() url.Values {
	q := make(url.Values, 4)
	if v := b.id.ID(); !v.Empty() {
		q.Set("id", v.String())
	}
	query.MergeInto(
		q,
		b.names.URLQuery(),
		b.pg.URLQuery(),
	)
	return q
}
