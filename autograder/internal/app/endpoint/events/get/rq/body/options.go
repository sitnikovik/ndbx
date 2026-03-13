package body

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/pagination"
)

// Option represents a functional option for configuring the Body struct.
type Option func(*Body)

// WithPagination sets the pagination for the Body struct.
func WithPagination(pg pagination.Pagination) Option {
	return func(b *Body) {
		b.pg = pg
	}
}

// WithTitle sets the event title for the Body struct.
func WithTitle(title string) Option {
	return func(b *Body) {
		b.content.title = title
	}
}
