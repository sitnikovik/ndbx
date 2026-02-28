package redis

import (
	"context"

	"github.com/sitnikovik/ndbx/autograder/internal/client/redis/valtype"
)

// Type retrieves the type of the value stored at the given key in Redis.
//
// It returns the type of the value, or an error if the operation fails.
func (c *Client) Type(
	ctx context.Context,
	key string,
) (valtype.Type, error) {
	t, err := c.cli.Type(ctx, key).Result()
	if err != nil {
		return valtype.None, err
	}
	return valtype.ParseType(t), nil
}
