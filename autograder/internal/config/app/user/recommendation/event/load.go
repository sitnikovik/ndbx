package event

import (
	"time"

	"github.com/sitnikovik/ndbx/autograder/internal/env"
)

// Load loads session configuration from environment variablesand returns it.
func Load() Config {
	ttl := env.
		Get("APP_RECOMMENDATIONS_TTL").
		Int()
	return NewConfig(
		time.Duration(ttl) * time.Second,
	)
}
