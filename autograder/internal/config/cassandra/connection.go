package cassandra

import (
	"errors"
)

// Connection defines a connection parameters to Apache Cassandra.
type Connection struct {
	// hosts is the list of Cassandra hosts.
	hosts []string
	// port is the Cassandra port.
	port int
}

// NewConnection creates a new Connection instance.
func NewConnection(
	hosts []string,
	port int,
) Connection {
	return Connection{
		hosts: hosts,
		port:  port,
	}
}

// Port returns the port number for the Cassandra cluster.
func (c Connection) Port() int {
	return c.port
}

// Hosts returns the list of Cassandra hosts we should connect to.
func (c Connection) Hosts() []string {
	return c.hosts
}

// Validate validates the connection config
// and returns if it has any invalid value.
func (c Connection) Validate() error {
	if len(c.Hosts()) == 0 {
		return errors.New("empty hosts")
	}
	if p := c.Port(); p == 0 || p > 99999 {
		return errors.New("invalid port")
	}
	return nil
}
