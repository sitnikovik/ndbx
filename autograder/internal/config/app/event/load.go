package event

import (
	reaction "github.com/sitnikovik/ndbx/autograder/internal/config/app/reaction/event"
	review "github.com/sitnikovik/ndbx/autograder/internal/config/app/review/event"
)

// Load loads the event configuration from the environment variables and returns it.
//
// No panics if any required variable is missing.
func Load() Config {
	return NewConfig(
		reaction.Load(),
		review.Load(),
	)
}
