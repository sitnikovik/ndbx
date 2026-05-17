package user

import (
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user/recommendation"
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user/session"
)

// Config represents the configuration for the user component of the application.
type Config struct {
	// session is the configuration for user sessions.
	session session.Config
	// recomms is the configuration for user recommendations.
	recomms recommendation.Config
}

// NewConfig creates a new Config with the specified session configuration
// and optional functional options.
func NewConfig(session session.Config, opts ...Option) Config {
	c := Config{
		session: session,
	}
	for _, o := range opts {
		o(&c)
	}
	return c
}

// Session returns the configuration for user sessions.
func (c Config) Session() session.Config {
	return c.session
}

// Recommendations returns the configuration for user recommendations.
func (c Config) Recommendations() recommendation.Config {
	return c.recomms
}
