package reaction

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	common "github.com/sitnikovik/ndbx/autograder/internal/app/reaction"
	eventreaction "github.com/sitnikovik/ndbx/autograder/internal/app/reaction/event"
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

// ToLikes converts the scanner to the likes and returns them.
func ToLikes(iter scanner) []eventreaction.Like {
	defer func() {
		_ = iter.Close()
	}()
	var (
		likes     = make([]eventreaction.Like, 0)
		evid      string
		val       int8
		createdAt time.Time
		createdBy string
	)
	for iter.Scan(
		&evid,
		&val,
		&createdAt,
		&createdBy,
	) {
		if val != 1 {
			continue
		}
		likes = append(likes, eventreaction.NewLike(
			common.NewLike(
				creation.NewStamp(
					creation.NewCreated(
						createdAt,
						user.NewIdentity(
							user.NewID(createdBy),
						),
					),
				),
			),
			eventreaction.NewEvent(
				event.NewID(evid),
			),
		))
	}
	return likes
}
