package neo4j

// Config holds the Neo4j database configuration.
type Config struct {
	// conn is the Neo4j database connection details.
	conn Connection
	// auth is the Neo4j database authentication details.
	auth Auth
}

// NewConfig creates a new Neo4j database configuration.
func NewConfig(
	conn Connection,
	auth Auth,
) Config {
	return Config{
		conn: conn,
		auth: auth,
	}
}

// Connection returns the Neo4j database connection details.
func (c Config) Connection() Connection {
	return c.conn
}

// Auth returns the Neo4j database authentication details.
func (c Config) Auth() Auth {
	return c.auth
}
