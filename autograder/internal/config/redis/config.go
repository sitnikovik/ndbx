package redis

// Config holds the configuration for connecting to a Redis server.
type Config struct {
	// addr is the Redis server address (host:port).
	addr string
	// password is the Redis password (if any).
	password string
	// db is the Redis database number.
	db int
}

// NewConfig creates a new Redis Config with the given parameters.
//
// Parameters:
//   - addr: The address of the Redis server (e.g., "localhost:6379").
//   - password: The password for the Redis server (if required).
//   - db: The database number to use (default is 0).
func NewConfig(
	addr string,
	password string,
	db int,
) Config {
	return Config{
		addr:     addr,
		password: password,
		db:       db,
	}
}

// Addr returns the Redis server address.
//
// This is the host and port of the Redis server (e.g., "localhost:6379").
func (c Config) Addr() string {
	return c.addr
}

// Password returns the Redis password.
func (c Config) Password() string {
	return c.password
}

// DB returns the Redis database number.
//
// Redis databases are numbered from 0 to 15 by default,
// but this can be configured on the server side.
//
// The default database is 0.
func (c Config) DB() int {
	return c.db
}
