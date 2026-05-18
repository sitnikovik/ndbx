package event

import "time"

// Config represents a configuration for the event recommendations.
type Config struct {
	// ttl is the time-to-live for the event recommendations in cache.
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

// TTL returns the time-to-live for event recommendations.
func (c Config) TTL() time.Duration {
	return c.ttl
}
