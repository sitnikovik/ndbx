//go:build integration
// +build integration

package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/teardown/redis"
	redisclient "github.com/sitnikovik/ndbx/autograder/internal/client/redis"
	cfg "github.com/sitnikovik/ndbx/autograder/internal/config/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

func TestRedisTeardownJob_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}
	t.Run("teardown job cleans redis and closes connection successfully", func(t *testing.T) {
		ctx := context.Background()
		client := redisclient.NewClient(
			cfg.NewConfig(
				"localhost:6379",
				"",
				1,
			),
		)
		defer func() {
			_ = client.Close(ctx)
		}()
		err := client.Set(
			ctx,
			"test-key",
			"test-value",
			5*time.Minute,
		)
		require.NoError(t, err)
		has, err := client.Has(ctx, "test-key")
		require.NoError(t, err)
		require.True(t, has, "test data should exist before teardown")
		err = redis.
			NewJob(client).
			Run(
				ctx,
				step.NewVariables(),
			)
		assert.NoError(t, err)
	})
	t.Run("teardown job cleans multiple databases", func(t *testing.T) {
		ctx := context.Background()
		client1 := redisclient.NewClient(
			cfg.NewConfig(
				"localhost:6379",
				"",
				1,
			),
		)
		defer func() {
			_ = client1.Close(ctx)
		}()
		client2 := redisclient.NewClient(
			cfg.NewConfig(
				"localhost:6379",
				"",
				2,
			),
		)
		defer func() {
			_ = client2.Close(ctx)
		}()
		err := client1.Set(ctx, "key-db1", "value1", 5*time.Minute)
		require.NoError(t, err)
		err = client2.Set(ctx, "key-db2", "value2", 5*time.Minute)
		require.NoError(t, err)
		has1, err := client1.Has(ctx, "key-db1")
		require.NoError(t, err)
		require.True(t, has1)
		has2, err := client2.Has(ctx, "key-db2")
		require.NoError(t, err)
		require.True(t, has2)
		err = redis.
			NewJob(client1).
			Run(
				ctx,
				step.NewVariables(),
			)
		assert.NoError(t, err)
	})
	t.Run("teardown job handles empty redis", func(t *testing.T) {
		ctx := context.Background()
		client := redisclient.NewClient(
			cfg.NewConfig(
				"localhost:6379",
				"",
				3,
			),
		)
		defer func() {
			_ = client.Close(ctx)
		}()
		err := client.FlushAll(ctx)
		require.NoError(t, err)
		err = redis.
			NewJob(client).
			Run(
				ctx,
				step.NewVariables(),
			)
		assert.NoError(t, err)
	})
}
