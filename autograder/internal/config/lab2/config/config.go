package config

import (
	"github.com/sitnikovik/ndbx/autograder/internal/config/app"
	"github.com/sitnikovik/ndbx/autograder/internal/config/redis"
)

// Config represents the configuration for the job component of the autograder application.
type Config struct {
	// redis is the configuration for connecting to the Redis server.
	redis redis.Config
	// app is the configuration for the application settings.
	app app.Config
}

// NewConfig creates a new Config instance
// with the provided Redis and application configurations.
func NewConfig(
	redis redis.Config,
	app app.Config,
) Config {
	return Config{
		redis: redis,
		app:   app,
	}
}

// Redis returns the Redis configuration from the Config instance.
func (c Config) Redis() redis.Config {
	return c.redis
}

// App returns the application configuration from the Config instance.
func (c Config) App() app.Config {
	return c.app
}
