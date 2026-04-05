package event

import reaction "github.com/sitnikovik/ndbx/autograder/internal/config/app/reaction/event"

// Config represents the configuration for the events.
type Config struct {
	// reactions is the configuration for the event reactions.
	reactions reaction.Config
}

// NewConfig creates a new Config instance.
func NewConfig(react reaction.Config) Config {
	return Config{
		reactions: react,
	}
}

// Reactions returns the configuration for the event reactions.
func (c Config) Reactions() reaction.Config {
	return c.reactions
}
