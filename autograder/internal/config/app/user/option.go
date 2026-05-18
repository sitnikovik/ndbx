package user

import (
	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user/recommendation"
)

// Option is a functional option for configuring the configuration.
type Option func(*Config)

// WithRecommendations sets the recommendation configuration
// for the config being built.
func WithRecommendations(cfg recommendation.Config) Option {
	return func(c *Config) {
		c.recomms = cfg
	}
}
