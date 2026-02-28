package session

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/env"
)

// MustLoad loads session configuration from environment variables
// and panics if any required variable is missing or invalid.
func MustLoad() Config {
	ttl := env.
		MustGet("APP_USER_SESSION_TTL").
		MustInt()
	return NewConfig(
		time.Duration(ttl) * time.Second,
	)
}
