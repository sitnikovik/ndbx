package body

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/pagination"
	rangeof "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/range-of"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/category"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
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

// WithCategory sets the event category for the Body struct.
func WithCategory(cat category.Type) Option {
	return func(b *Body) {
		b.content.cat = cat
	}
}

// WithAddress sets the event address for the Body struct.
func WithAddress(addr string) Option {
	return func(b *Body) {
		b.loc.address = addr
	}
}

// WithCity sets the event city for the Body struct.
func WithCity(city string) Option {
	return func(b *Body) {
		b.loc.city = city
	}
}

// WithEntryPrice sets the event price for the Body struct.
func WithEntryPrice(from, to uint64) Option {
	return func(b *Body) {
		b.costs.entry = rangeof.NewUInts(
			"price",
			rangeof.NewUInt(from),
			rangeof.NewAnyUInt(to),
		)
	}
}

// WithDates sets the event date for the Body struct.
func WithDates(from, to time.Time) Option {
	return func(b *Body) {
		b.created.at = rangeof.NewDates(
			"date",
			from,
			to,
		)
	}
}

// WithByUser sets the user who created the event.
func WithByUser(usr user.Identity) Option {
	return func(b *Body) {
		b.created.by = usr
	}
}
