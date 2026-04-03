package cassandra

import (
	"github.com/gocql/gocql"

	config "github.com/sitnikovik/ndbx/autograder/internal/config/cassandra"
)

// Client represents a Cassandra client to interact with the database.
type Client struct {
	// cluster is the underlying Cassandra cluster we connect to.
	cluster *gocql.Session
	// cfg is the configuration for the Cassandra client.
	cfg config.Config
}

// NewClient creates a new Cassandra client
// using the provided configuration.
func NewClient(cfg config.Config) *Client {
	return &Client{
		cfg: cfg,
	}
}

// MustConnect connects to the Cassandra database
// and panics if the connection fails.
func (c *Client) MustConnect() {
	if err := c.Connect(); err != nil {
		panic(err)
	}
}

// Connect connects to the Cassandra database
// and returns an error if the connection fails.
func (c *Client) Connect() error {
	cluster := gocql.NewCluster(
		c.cfg.
			Connection().
			Hosts()...,
	)
	cluster.Port = c.cfg.
		Connection().
		Port()
	cluster.Keyspace = c.cfg.
		Database().
		Keyspace()
	cluster.Consistency = c.cfg.
		Database().
		Consistency().
		ToCQL()
	if auth := c.cfg.Auth(); !auth.Empty() {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: auth.Username(),
			Password: auth.Password(),
		}
	}
	session, err := cluster.CreateSession()
	if err != nil {
		return err
	}
	c.cluster = session
	return nil
}

// Close closes the connection to the Cassandra database.
func (c *Client) Close() error {
	c.cluster.Close()
	return nil
}
