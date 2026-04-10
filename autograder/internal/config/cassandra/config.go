package cassandra

// Config defines Cassandra configuration.
type Config struct {
	// conn is the Cassandra connection parameters.
	conn Connection
	// usr is the username for Cassandra authentication.
	auth Auth
	// db is the Cassandra database queries configuration.
	db Database
}

// NewConfig creates a new Cassandra configuration with the provided parameters.
func NewConfig(
	conn Connection,
	auth Auth,
	db Database,
) Config {
	return Config{
		conn: conn,
		auth: auth,
		db:   db,
	}
}

// Connection returns the connection parameters for Cassandra.
func (c Config) Connection() Connection {
	return c.conn
}

// Auth returns the authentication configuration for Cassandra.
func (c Config) Auth() Auth {
	return c.auth
}

// Database returns the database configuration for Cassandra.
func (c Config) Database() Database {
	return c.db
}
