package event

import "github.com/sitnikovik/ndbx/autograder/internal/config/app/reaction/event"

// Load loads the event configuration from the environment variables and returns it.
//
// No panics if any required variable is missing.
func Load() Config {
	return NewConfig(
		event.Load(),
	)
}
