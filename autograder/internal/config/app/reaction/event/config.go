package event

import "time"

// Config represents a configuration for the event reactions.
type Config struct {
	// ttl is duration for event reactions in cache.
	ttl time.Duration
}

// NewConfig creates a new Config instance.
func NewConfig(
	ttl time.Duration,
) Config {
	return Config{
		ttl: ttl,
	}
}

// TTL returns the time-to-live for event reactions.
func (c Config) TTL() time.Duration {
	return c.ttl
}
