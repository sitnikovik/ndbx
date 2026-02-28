package app

import (
	"strconv"

	"github.com/sitnikovik/ndbx/autograder/internal/config/app/user"
)

// Config is the configuration for the application.
type Config struct {
	// user is the configuration for user management.
	user user.Config
	// host is the host address for the application.
	host string
	// port is the port number for the application.
	port int
}

// NewConfig creates a new Config with the given parameters.
//
// Parameters:
//   - user: The configuration for user management.
//   - host: The host address for the application.
//   - port: The port number for the application.
func NewConfig(
	user user.Config,
	host string,
	port int,
) Config {
	return Config{
		user: user,
		host: host,
		port: port,
	}
}

// User returns the configuration for user management.
func (c Config) User() user.Config {
	return c.user
}

// Host returns the host address for the application.
func (c Config) Host() string {
	return c.host
}

// Port returns the port number for the application.
func (c Config) Port() int {
	return c.port
}

// Address returns the full address (host:port) for the application.
func (c Config) Address() string {
	return c.host + ":" + strconv.Itoa(c.port)
}
