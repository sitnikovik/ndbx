package config

import (
	"github.com/sitnikovik/ndbx/autograder/internal/config/app"
	"github.com/sitnikovik/ndbx/autograder/internal/config/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/config/neo4j"
)

// Config represents the configuration for the job component of the autograder application.
type Config struct {
	// mongo is the configuration for connecting to the MongoDB server.
	mongo mongo.Config
	// neo4j is the configuration for connecting to the Neo4J server.
	neo4j neo4j.Config
	// app is the configuration for the application settings.
	app app.Config
}

// NewConfig creates a new Config instance
// with the provided Redis, MongoDB, Neo4J and application configurations.
func NewConfig(
	mongo mongo.Config,
	neo4j neo4j.Config,
	app app.Config,
) Config {
	return Config{
		mongo: mongo,
		neo4j: neo4j,
		app:   app,
	}
}

// Mongo returns the MongoDB configuration from the Config instance.
func (c Config) Mongo() mongo.Config {
	return c.mongo
}

// Neo4j returns the Neo4J configuration from the Config instance.
func (c Config) Neo4j() neo4j.Config {
	return c.neo4j
}

// App returns the application configuration from the Config instance.
func (c Config) App() app.Config {
	return c.app
}
