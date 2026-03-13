package mongo

import "fmt"

// Config holds the MongoDB connection configuration.
type Config struct {
	// db is the name of the MongoDB database.
	db string
	// usr is the username for MongoDB authentication.
	usr string
	// pwd is the password for MongoDB authentication.
	pwd string
	// host is the MongoDB server host.
	host string
	// port is the MongoDB server port.
	port int
}

// NewConfig creates a new MongoDB configuration with the provided parameters.
//
// Parameters:
//   - db: The name of the MongoDB database.
//   - usr: The username for MongoDB authentication.
//   - pwd: The password for MongoDB authentication.
//   - host: The MongoDB server host.
//   - port: The MongoDB server port.
func NewConfig(
	db string,
	usr string,
	pwd string,
	host string,
	port int,
) Config {
	return Config{
		db:   db,
		usr:  usr,
		pwd:  pwd,
		host: host,
		port: port,
	}
}

// URI returns the MongoDB connection URI.
func (c Config) URI() string {
	// Если credentials пустые, подключаемся без аутентификации
	if c.usr == "" && c.pwd == "" {
		return fmt.Sprintf(
			"mongodb://%s:%d/%s",
			c.host,
			c.port,
			c.db,
		)
	}
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/%s",
		c.usr,
		c.pwd,
		c.host,
		c.port,
		c.db,
	)
}

// Database returns the name of the MongoDB database to use.
func (c Config) Database() string {
	return c.db
}
