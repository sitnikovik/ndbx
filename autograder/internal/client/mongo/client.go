package mongo

import (
	"context"

	drv "go.mongodb.org/mongo-driver/v2/mongo"
	drvopts "go.mongodb.org/mongo-driver/v2/mongo/options"

	config "github.com/sitnikovik/ndbx/autograder/internal/config/mongo"
)

// Client is a wrapper around the MongoDB client
// that provides methods for interacting with the MongoDB server.
type Client struct {
	// cli is the underlying MongoDB client.
	cli *drv.Client
	// uri is the MongoDB connection URI.
	uri string
	// db is the name of the database to use.
	db string
	// connected indicates whether the client is currently connected to the MongoDB server.
	connected bool
}

// NewClient creates a new MongoDB client using the provided configuration.
func NewClient(cfg config.Config) *Client {
	return &Client{
		uri: cfg.URI(),
		db:  cfg.Database(),
	}
}

// Connect establishes a connection to the MongoDB server using the provided URI.
//
// It should be called before using the client to interact with the MongoDB server.
// Returns an error if the connection fails.
func (c *Client) Connect() error {
	cli, err := drv.Connect(
		drvopts.
			Client().
			ApplyURI(c.uri).
			SetServerAPIOptions(
				drvopts.ServerAPI(drvopts.ServerAPIVersion1),
			),
	)
	if err != nil {
		return err
	}
	c.cli = cli
	c.connected = true
	return nil
}

// MustConnect establishes a connection to the MongoDB server and panics if the connection fails.
func (c *Client) MustConnect() {
	if err := c.Connect(); err != nil {
		panic(err)
	}
}

// Close closes the connection to the MongoDB server.
//
// It should be called when the client is no longer needed to free up resources.
// Returns an error if the operation fails.
func (c *Client) Close(ctx context.Context) error {
	return c.cli.Disconnect(ctx)
}
