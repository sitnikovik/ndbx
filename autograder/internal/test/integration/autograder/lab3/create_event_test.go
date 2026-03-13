//go:build integration
// +build integration

package lab3_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	eventdox "github.com/sitnikovik/ndbx/autograder/internal/app/mongo/event/doc"
	"github.com/sitnikovik/ndbx/autograder/internal/app/mongo/event/doc/key"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
	config "github.com/sitnikovik/ndbx/autograder/internal/config/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/env"
)

func TestClient_CreateEvent(t *testing.T) {
	cli := mongo.NewClient(
		config.NewConfig(
			env.MustGet("MONGODB_DATABASE").String(),
			env.Get("MONGODB_USER").String(),
			env.Get("MONGODB_PASSWORD").String(),
			env.MustGet("MONGODB_HOST").String(),
			env.MustGet("MONGODB_PORT").MustInt(),
		),
	)
	t.Cleanup(func() {
		err := cli.Close(context.Background())
		require.NoError(t, err)
	})
	t.Run("insert and get", func(t *testing.T) {
		var err error
		ctx := context.Background()
		err = cli.InsertOne(
			ctx,
			"test_collection",
			doc.NewKVs(
				doc.NewKV(
					key.Title,
					"EventHub Fest",
				),
				doc.NewKV(
					key.Description,
					"The best showtime ever...",
				),
				doc.NewKV(
					key.Location,
					"Main Street 13",
				),
				doc.NewKV(
					key.CreatedAt,
					"2025-01-01T12:00:00Z",
				),
				doc.NewKV(
					key.CreatedBy,
					"123",
				),
				doc.NewKV(
					key.StartedAt,
					"2025-01-02T12:00:00Z",
				),
				doc.NewKV(
					key.FinishedAt,
					"2025-01-02T14:00:00Z",
				),
			),
		)
		require.NoError(t, err)
		all, err := cli.AllBy(
			ctx,
			"test_collection",
			doc.NewKVs(
				doc.NewKV(
					key.Title,
					"EventHub Fest",
				),
			),
		)
		require.NoError(t, err)
		id := all.First().ID()
		require.NotEmpty(t, id)
		evnt, err := cli.ByID(
			ctx,
			"test_collection",
			id,
		)
		require.NoError(t, err)
		got := eventdox.
			NewEventDocument(evnt).
			ToEvent()
		assert.Equal(
			t,
			"EventHub Fest",
			got.
				Content().
				Title(),
		)
		assert.Equal(
			t,
			"The best showtime ever...",
			got.
				Content().
				Description(),
		)
		assert.Equal(
			t,
			"Main Street 13",
			got.
				Location().
				Address(),
		)
		assert.Equal(
			t,
			"2025-01-01T12:00:00Z",
			got.
				Created().
				At().
				Format(time.RFC3339),
		)
		assert.Equal(
			t,
			"123",
			got.
				Created().
				By().
				ID().
				String(),
		)
		assert.Equal(
			t,
			"2025-01-02T12:00:00Z",
			got.
				Dates().
				StartedAt().
				Format(time.RFC3339),
		)
		assert.Equal(
			t,
			"2025-01-02T14:00:00Z",
			got.
				Dates().
				FinishedAt().
				Format(time.RFC3339),
		)
	})
}
