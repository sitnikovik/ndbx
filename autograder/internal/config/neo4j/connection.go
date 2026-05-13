package neo4j

import "errors"

// Connection is the Neo4j database connection configuration.
type Connection struct {
	// url is the Neo4j database URL
	url string
}

// NewConnection creates a new Neo4j connection configuration instance.
func NewConnection(u string) Connection {
	return Connection{
		url: u,
	}
}

// URL returns the Neo4j database URL.
func (c Connection) URL() string {
	return c.url
}

// Validate validates the connection configuration.
//
// Not validates the URL format so if the URL is invalid, nil is returned.
func (c Connection) Validate() error {
	if c.url == "" {
		return errors.New("empty url")
	}
	return nil
}
