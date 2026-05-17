package recommendation

import "github.com/sitnikovik/ndbx/autograder/internal/config/app/user/recommendation/event"

// Config is the configuration for the recommendation.
type Config struct {
	// events is the configuration for the events.
	events event.Config
}

// NewConfig creates a new Config instance with the given options.
func NewConfig(opt ...Option) Config {
	c := Config{}
	for _, o := range opt {
		o(&c)
	}
	return c
}

// Events returns the configuration for the events.
func (c Config) Events() event.Config {
	return c.events
}
