package user

import "github.com/sitnikovik/ndbx/autograder/internal/config/app/user/session"

// MustLoad loads the user configuration
// and panics if any required configuration is missing or invalid.
func MustLoad() Config {
	return Config{
		session: session.MustLoad(),
	}
}
