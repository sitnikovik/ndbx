package config

import (
	"github.com/sitnikovik/ndbx/autograder/internal/config/app"
	"github.com/sitnikovik/ndbx/autograder/internal/config/cassandra"
	"github.com/sitnikovik/ndbx/autograder/internal/config/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/config/neo4j"
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
	// neo4j is the configuration for connecting to the Neo4J server.
	neo4j neo4j.Config
	// app is the configuration for the application settings.
	app app.Config
}

// NewConfig creates a new Config instance
// with the provided Redis, MongoDB, Neo4J and application configurations.
func NewConfig(
	redis redis.Config,
	mongo mongo.Config,
	cassandra cassandra.Config,
	neo4j neo4j.Config,
	app app.Config,
) Config {
	return Config{
		redis:     redis,
		mongo:     mongo,
		cassandra: cassandra,
		neo4j:     neo4j,
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

// Neo4j returns the Neo4J configuration from the Config instance.
func (c Config) Neo4j() neo4j.Config {
	return c.neo4j
}

// App returns the application configuration from the Config instance.
func (c Config) App() app.Config {
	return c.app
}
