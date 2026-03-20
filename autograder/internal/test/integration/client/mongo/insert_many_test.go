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

func TestClient_Insert(t *testing.T) {
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
	t.Run("insert many", func(t *testing.T) {
		var err error
		ctx := context.Background()
		err = cli.Insert(
			ctx,
			"test_collection",
			doc.NewKVs(
				doc.NewKV("name", "Frodo Baggings"),
				doc.NewKV("username", "fr0dothel0tr"),
			),
			doc.NewKVs(
				doc.NewKV("name", "John Doe"),
				doc.NewKV("username", "j0hnd0e42"),
			),
		)
		require.NoError(t, err)
	})
}
