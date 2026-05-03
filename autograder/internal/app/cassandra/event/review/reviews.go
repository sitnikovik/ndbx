package review

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra"
	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra/event/review/filter"
	eventreview "github.com/sitnikovik/ndbx/autograder/internal/app/review/event"
	qb "github.com/sitnikovik/ndbx/autograder/internal/client/cassandra/query/builder"
)

// Reviews represents event reviews from database.
type Reviews struct {
	// db is database client.
	db cassandra.Selectable
	// ftr is filter for reactions.
	ftr *filter.Filter
	// limit defines limit of returned reviews.
	limit int
}

// NewReviews returns new Reviews from database.
func NewReviews(
	db cassandra.Selectable,
	opts ...Option,
) *Reviews {
	r := &Reviews{
		db: db,
		ftr: filter.NewFilter(
			qb.NewWhere(),
		),
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// Select returns all reviews for the event.
func (r *Reviews) Select(ctx context.Context) ([]eventreview.Review, error) {
	q := fmt.Sprintf(
		`SELECT %s, %s, %s, %s, %s, %s, %s FROM %s`,
		"id",
		"event_id",
		"rating",
		"comment",
		"created_at",
		"created_by",
		"updated_at",
		Table,
	)
	if !r.ftr.Empty() {
		q += " " + r.ftr.Where()
	}
	if n := r.limit; n > 0 {
		q += ` LIMIT ` + strconv.Itoa(n)
	}
	var itr cassandra.Scanner
	var err error
	if !r.ftr.Empty() {
		itr, err = r.db.Select(ctx, q, r.ftr.Args()...)
	} else {
		itr, err = r.db.Select(ctx, q)
	}
	if err != nil {
		return nil, err
	}
	return ToReviews(itr), nil
}
