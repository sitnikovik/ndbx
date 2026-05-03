package review

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/rating"
	eventreview "github.com/sitnikovik/ndbx/autograder/internal/app/review/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

// scanner implements the scanner
// that can be used to scan the likes.
type scanner interface {
	// Scan scans the next like and writes
	// the values to the provided variables.
	Scan(...any) bool
	// Close closes the scanner.
	Close() error
}

// ToReviews converts the scanner to event reviews and returns them.
func ToReviews(iter scanner) []eventreview.Review {
	defer func() {
		_ = iter.Close()
	}()
	var (
		reviews   = make([]eventreview.Review, 0)
		createdAt time.Time
		updatedAt time.Time
		id        string
		evid      string
		createdBy string
		comment   string
		rate      int8
	)
	for iter.Scan(
		&id,
		&evid,
		&rate,
		&comment,
		&createdAt,
		&createdBy,
		&updatedAt,
	) {
		reviews = append(reviews, eventreview.NewReview(
			id,
			creation.NewStamp(
				creation.NewCreated(
					createdAt,
					user.NewIdentity(
						user.NewID(createdBy),
					),
				),
			),
			eventreview.NewEvent(
				event.NewID(evid),
			),
			comment,
			rating.NewRating(rate),
			eventreview.WithUpdatedAt(updatedAt),
		))
	}
	return reviews
}
