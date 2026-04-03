package cassandra

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
