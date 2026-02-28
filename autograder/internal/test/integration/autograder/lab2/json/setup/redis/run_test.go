//go:build integration
// +build integration

package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/setup/redis"
	redisclient "github.com/sitnikovik/ndbx/autograder/internal/client/redis"
	cfg "github.com/sitnikovik/ndbx/autograder/internal/config/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
)

func TestRedisSetupJob_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}
	client := redisclient.NewClient(
		cfg.NewConfig(
			"localhost:6379",
			"",
			1,
		),
	)
	ctx := context.Background()
	t.Cleanup(func() {
		_ = client.FlushAll(ctx)
		_ = client.Close(ctx)
	})
	t.Run("setup job cleans redis successfully", func(t *testing.T) {
		ctx := context.Background()
		err := client.Set(
			ctx,
			"test-key",
			"test-value",
			5*time.Minute,
		)
		require.NoError(t, err)
		has, err := client.Has(ctx, "test-key")
		require.NoError(t, err)
		require.True(t, has)
		err = redis.
			NewJob(client).
			Run(
				ctx,
				step.NewVariables(),
			)
		assert.NoError(t, err)
		empty, err := client.Empty(ctx)
		assert.NoError(t, err)
		assert.True(t, empty, "Redis should be empty after setup job")
	})
}
