package app

import "github.com/sitnikovik/ndbx/autograder/internal/config/app/event"

// Option is a functional option for configuring the app configuration.
type Option func(*Config)

// WithEvent sets the event configuration for the app configuration.
func WithEvent(event event.Config) Option {
	return func(c *Config) {
		c.event = event
	}
}
