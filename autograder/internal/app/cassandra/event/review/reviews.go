package review

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cassandra"
	eventreview "github.com/sitnikovik/ndbx/autograder/internal/app/review/event"
)

// Reviews represents event reviews from database.
type Reviews struct {
	// db is database client.
	db cassandra.Selectable
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
	if n := r.limit; n > 0 {
		q += ` LIMIT ` + strconv.Itoa(n)
	}
	itr, err := r.db.Select(ctx, q)
	if err != nil {
		return nil, err
	}
	return ToReviews(itr), nil
}
