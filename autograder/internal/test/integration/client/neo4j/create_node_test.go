package neo4j_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sitnikovik/ndbx/autograder/internal/client/neo4j"
	config "github.com/sitnikovik/ndbx/autograder/internal/config/neo4j"
)

func TestClient_CreateNode(t *testing.T) {
	cli := neo4j.NewClient(
		config.MustLoad(),
	)
	t.Cleanup(func() {
		err := cli.Close(context.Background())
		require.NoError(t, err)
	})
	ctx := context.Background()
	err := cli.ExecQuery(
		ctx,
		"CREATE (n:Item { id: $id, name: $name }) RETURN n",
		map[string]any{
			"id":   1,
			"name": "Item 1",
		},
	)
	require.NoError(
		t,
		err,
		"unexpected error: %v",
		err,
	)
}
