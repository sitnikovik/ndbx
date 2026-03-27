package body

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/pagination"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

// Option represents a functional option
// for configuring the Body instance.
type Option func(b *Body)

// WithFullName sets full name to the Names instance.
func WithFullName(s string) Option {
	return func(b *Body) {
		b.names.fullName = s
	}
}

// WithIdentity sets user identity to the Body instance.
func WithIdentity(id user.Identity) Option {
	return func(b *Body) {
		b.id = id
	}
}

// WithPagination sets the pagination for the Body instance.
func WithPagination(pg pagination.Pagination) Option {
	return func(b *Body) {
		b.pg = pg
	}
}
