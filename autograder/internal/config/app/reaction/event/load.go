package event

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/env"
)

// Load loads event configuration from environment variables and returns it.
//
// No panics if any required variable is missing.
func Load() Config {
	ttl := env.
		Get("APP_LIKE_TTL").
		MustInt()
	return NewConfig(
		time.Duration(ttl) * time.Second,
	)
}
