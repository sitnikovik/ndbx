package mongo

import "github.com/sitnikovik/ndbx/autograder/internal/env"

// MustLoad loads the MongoDB configuration from environment variables
// and panics if any required variable is missing or invalid.
func MustLoad() Config {
	return NewConfig(
		env.MustGet("MONGODB_DATABASE").String(),
		env.Get("MONGODB_USER").String(),
		env.Get("MONGODB_PASSWORD").String(),
		env.MustGet("MONGODB_HOST").String(),
		env.MustGet("MONGODB_PORT").MustInt(),
	)
}
