package redis

import (
	"fmt"

	"github.com/sitnikovik/ndbx/autograder/internal/env"
)

// MustLoad loads the Redis configuration from environment variables
// and panics if any required variable is missing or invalid.
func MustLoad() Config {
	return NewConfig(
		fmt.Sprintf(
			"%s:%s",
			env.MustGet("REDIS_HOST").String(),
			env.MustGet("REDIS_PORT").String(),
		),
		env.Get("REDIS_PASSWORD").String(),
		env.MustGet("REDIS_DB").MustInt(),
	)
}
