package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/app/neo4j/node/user"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

func TestUsers_OneWithID(t *testing.T) {
	t.Parallel()
	type args struct {
		id user.ID
	}
	type want struct {
		user impl.User
	}
	tests := []struct {
		name string
		uu   impl.Users
		args args
		want want
	}{
		{
			name: "found in list",
			uu: impl.NewUsers(
				impl.NewUser(
					user.NewID("123"),
				),
				impl.NewUser(
					user.NewID("213"),
				),
			),
			args: args{
				id: user.NewID("123"),
			},
			want: want{
				user: impl.NewUser(
					user.NewID("123"),
				),
			},
		},
		{
			name: "not found",
			uu: impl.NewUsers(
				impl.NewUser(
					user.NewID("213"),
				),
			),
			args: args{
				id: user.NewID("123"),
			},
			want: want{
				user: impl.User{},
			},
		},
		{
			name: "empty list",
			uu:   impl.NewUsers(),
			args: args{
				id: user.NewID("123"),
			},
			want: want{
				user: impl.User{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.user,
				tt.uu.OneWithID(tt.args.id),
			)
		})
	}
}
