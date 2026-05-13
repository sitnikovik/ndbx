package neo4j

import (
	"context"

	driver "github.com/neo4j/neo4j-go-driver/v6/neo4j"
	config "github.com/sitnikovik/ndbx/autograder/internal/config/neo4j"
)

// Client is a Neo4j client that provides methods
// for interacting with the database.
type Client struct {
	// cli is the Neo4j driver client.
	cli driver.Driver
	// cfg is the Neo4j configuration.
	cfg config.Config
}

// NewClient creates a new Neo4j client with the given configuration.
func NewClient(cfg config.Config) *Client {
	return &Client{
		cfg: cfg,
	}
}

// Connect connects to the Neo4j database.
func (c *Client) Connect() error {
	if c.cli != nil {
		return nil
	}
	drv, err := driver.NewDriver(
		c.cfg.Connection().URL(),
		driver.BasicAuth(
			c.cfg.Auth().Username(),
			c.cfg.Auth().Password(),
			"",
		),
	)
	if err != nil {
		return err
	}
	c.cli = drv
	return nil
}

// MustConnect connects to the Neo4j database
// and panics if an error occurs.
func (c *Client) MustConnect() {
	err := c.Connect()
	if err != nil {
		panic(err)
	}
}

// Close closes the Neo4j driver client.
func (c *Client) Close(ctx context.Context) error {
	return c.cli.Close(ctx)
}
