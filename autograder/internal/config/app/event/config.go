package event

import (
	reaction "github.com/sitnikovik/ndbx/autograder/internal/config/app/reaction/event"
	review "github.com/sitnikovik/ndbx/autograder/internal/config/app/review/event"
)

// Config represents the configuration for the events.
type Config struct {
	// reactions is the configuration for the event reactions.
	reactions reaction.Config
	// reviews is the configuration for the event reviews.
	reviews review.Config
}

// NewConfig creates a new Config instance.
func NewConfig(
	react reaction.Config,
	rev review.Config,
) Config {
	return Config{
		reactions: react,
		reviews:   rev,
	}
}

// Reactions returns the configuration for the event reactions.
func (c Config) Reactions() reaction.Config {
	return c.reactions
}

// Reviews returns the configuration for the event reviews.
func (c Config) Reviews() review.Config {
	return c.reviews
}
