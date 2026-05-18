package user_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	node "github.com/sitnikovik/ndbx/autograder/internal/app/neo4j/node/user"
	impl "github.com/sitnikovik/ndbx/autograder/internal/app/neo4j/user"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/client/neo4j/graph"
	neo4jfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/neo4j/client"
)

func TestUsers_All(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx   context.Context
		limit int
	}
	type want struct {
		users   node.Users
		errored bool
	}
	tests := []struct {
		name string
		u    *impl.Users
		args args
		want want
	}{
		{
			name: "ok",
			u: impl.NewUsers(
				neo4jfk.NewClient(
					neo4jfk.WithQueryNodes(
						func(
							_ context.Context,
							_ string,
							_ map[string]any,
							_ ...string,
						) (graph.Nodes, error) {
							return graph.NewNodes(
								graph.NewNode(
									"1",
									graph.PropertiesFromMap(
										map[string]any{
											"id": "53e9c0c1a2a3c3d7e6c9c8a1",
										},
									),
								),
							), nil
						},
					),
				),
			),
			args: args{
				ctx:   context.Background(),
				limit: 0,
			},
			want: want{
				users: node.NewUsers(
					node.NewUser(
						user.NewID("53e9c0c1a2a3c3d7e6c9c8a1"),
					),
				),
				errored: false,
			},
		},
		{
			name: "empty users",
			u: impl.NewUsers(
				neo4jfk.NewClient(
					neo4jfk.WithQueryNodes(
						func(
							_ context.Context,
							_ string,
							_ map[string]any,
							_ ...string,
						) (graph.Nodes, error) {
							return graph.NewNodes(), nil
						},
					),
				),
			),
			args: args{
				ctx:   context.Background(),
				limit: 1,
			},
			want: want{
				users:   node.Users{},
				errored: false,
			},
		},
		{
			name: "db failed",
			u: impl.NewUsers(
				neo4jfk.NewClient(
					neo4jfk.WithQueryNodes(
						func(
							_ context.Context,
							_ string,
							_ map[string]any,
							_ ...string,
						) (graph.Nodes, error) {
							return nil, assert.AnError
						},
					),
				),
			),
			args: args{
				ctx:   context.Background(),
				limit: 1,
			},
			want: want{
				users:   nil,
				errored: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.u.All(
				tt.args.ctx,
				tt.args.limit,
			)
			assert.Equal(
				t,
				tt.want.users,
				got,
			)
			if tt.want.errored {
				assert.Error(t, err)
			} else {
				assert.NoErrorf(
					t,
					err,
					"unexpected error: %v",
					err,
				)
			}
		})
	}
}
