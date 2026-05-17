package user

import (
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user/recommendation"
)

// Option is a functional option for configuring the configuration.
type Option func(*Config)

// WithEvent sets the event recommendation configuration
// to the instance of its creation.
func WithRecommendations(cfg recommendation.Config) Option {
	return func(c *Config) {
		c.recomms = cfg
	}
}
