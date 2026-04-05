package app

import (
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/env"
)

// MustLoad loads the application configuration
// and panics if any required configuration is missing or invalid.
func MustLoad() Config {
	return NewConfig(
		user.MustLoad(),
		env.MustGet("APP_HOST").String(),
		env.MustGet("APP_PORT").MustInt(),
		WithEvent(event.Load()),
	)
}
