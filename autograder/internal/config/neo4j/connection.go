package neo4j

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
