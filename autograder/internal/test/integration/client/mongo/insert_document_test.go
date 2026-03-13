//go:build integration
// +build integration

package mongo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/client/mongo/doc"
	config "github.com/sitnikovik/ndbx/autograder/internal/config/mongo"
	"github.com/sitnikovik/ndbx/autograder/internal/env"
)

func TestClient_InsertOne(t *testing.T) {
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
	t.Run("insert and find", func(t *testing.T) {
		var err error
		ctx := context.Background()
		err = cli.InsertOne(
			ctx,
			"test_collection",
			doc.NewKVs(
				doc.NewKV("name", "Sam Sepiol"),
				doc.NewKV("username", "sams3piol"),
			),
		)
		require.NoError(t, err)
		ff, err := cli.AllBy(
			ctx,
			"test_collection",
			doc.NewKVs(
				doc.NewKV("name", "Sam Sepiol"),
			),
		)
		require.NoError(t, err)
		require.Len(t, ff, 1)
		require.Equal(
			t,
			"Sam Sepiol",
			ff.First().
				KVs().
				First().
				Value(),
		)
		require.Equal(
			t,
			"sams3piol",
			ff.First().
				KVs().
				Last().
				Value(),
		)
	})
}
