package user

import "github.com/sitnikovik/ndbx/autograder/internal/config/app/user/session"

// Config represents the configuration for the user component of the application.
type Config struct {
	// session is the configuration for user sessions.
	session session.Config
}

// NewConfig creates a new Config with the specified session configuration.
func NewConfig(session session.Config) Config {
	return Config{
		session: session,
	}
}

// Session returns the configuration for user sessions.
func (c Config) Session() session.Config {
	return c.session
}
