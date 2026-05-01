package filter

import (
	"github.com/sitnikovik/ndbx/autograder/internal/app/creation"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

// Option represents a functional option
// to configure filtering the Reactions instance.
type Option func(f *Filter)

// WithEventID sets the event ID
// to filter the reactions by that in the database.
func WithEventID(id event.ID) Option {
	return func(f *Filter) {
		f.eventID = id
	}
}

// WithCreatedBy sets the user's ID
// to filter the reactions by that in the database.
func WithCreatedBy(id user.ID) Option {
	return func(f *Filter) {
		f.stamp = creation.NewStamp(
			creation.NewCreated(
				f.stamp.Created().At(),
				user.NewIdentity(id),
			),
		)
	}
}
