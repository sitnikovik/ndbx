package pagination

import (
	"net/url"
	"strconv"
)

// Pagination represents the pagination parameters for listing items in an endpoint.
type Pagination struct {
	// limit is the maximum number of items to return.
	limit uint64
	// offset is the number of items to skip before starting to collect the result set.
	offset uint64
}

// NewPagination creates a new Pagination instance.
func NewPagination(limit, offset uint64) Pagination {
	return Pagination{
		limit:  limit,
		offset: offset,
	}
}

// URLQuery converts the Pagination into url.Values.
func (p Pagination) URLQuery() url.Values {
	q := make(url.Values, 2)
	if v := p.limit; v != 0 {
		q.Set("limit", strconv.FormatUint(v, 10))
	}
	if v := p.offset; v != 0 {
		q.Set("offset", strconv.FormatUint(v, 10))
	}
	return q
}
