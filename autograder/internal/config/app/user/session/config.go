package session

import "time"

// Config holds the configuration for user sessions.
type Config struct {
	// ttl is the time-to-live for user sessions.
	ttl time.Duration
}

// NewConfig creates a new Config with the given parameters.
//
// Parameters:
//   - ttl: The time-to-live for user sessions.
func NewConfig(
	ttl time.Duration,
) Config {
	return Config{
		ttl: ttl,
	}
}

// TTL returns the time-to-live for user sessions.
func (c Config) TTL() time.Duration {
	return c.ttl
}
