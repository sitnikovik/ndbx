package mongo

import "context"

// Ping checks the connection to the MongoDB server.
func (c *Client) Ping(ctx context.Context) error {
	if !c.connected {
		err := c.Connect()
		if err != nil {
			return err
		}
	}
	return c.cli.Ping(ctx, nil)
}
