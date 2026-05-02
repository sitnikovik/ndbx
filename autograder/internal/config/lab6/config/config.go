package config

import (
	"github.com/sitnikovik/ndbx/autograder/internal/config/app"
	"github.com/sitnikovik/ndbx/autograder/internal/config/cassandra"
	"github.com/sitnikovik/ndbx/autograder/internal/config/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/config/redis"
)

// Config represents the configuration for the job component of the autograder application.
type Config struct {
	// redis is the configuration for connecting to the Redis server.
	redis redis.Config
	// mongo is the configuration for connecting to the MongoDB server.
	mongo mongo.Config
	// cassandra is the configuration for connecting to the Apache Cassandra server.
	cassandra cassandra.Config
	// app is the configuration for the application settings.
	app app.Config
}

// NewConfig creates a new Config instance
// with the provided Redis and application configurations.
func NewConfig(
	redis redis.Config,
	mongo mongo.Config,
	cassandra cassandra.Config,
	app app.Config,
) Config {
	return Config{
		redis:     redis,
		mongo:     mongo,
		cassandra: cassandra,
		app:       app,
	}
}

// Redis returns the Redis configuration from the Config instance.
func (c Config) Redis() redis.Config {
	return c.redis
}

// Mongo returns the MongoDB configuration from the Config instance.
func (c Config) Mongo() mongo.Config {
	return c.mongo
}

// Cassandra returns the Apache Cassandra configuration from the Config instance.
func (c Config) Cassandra() cassandra.Config {
	return c.cassandra
}

// App returns the application configuration from the Config instance.
func (c Config) App() app.Config {
	return c.app
}
