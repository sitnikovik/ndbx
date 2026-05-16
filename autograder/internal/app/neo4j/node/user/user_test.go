package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/app/neo4j/node/user"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

func TestUser_ID(t *testing.T) {
	t.Parallel()
	type want struct {
		value user.ID
	}
	tests := []struct {
		name string
		u    impl.User
		want want
	}{
		{
			name: "ok",
			u: impl.NewUser(
				user.NewID("65e9c0b1a2b3c4d5e6f7a8b9"),
			),
			want: want{
				value: user.NewID("65e9c0b1a2b3c4d5e6f7a8b9"),
			},
		},
		{
			name: "empty id",
			u: impl.NewUser(
				user.NewID(""),
			),
			want: want{
				value: user.NewID(""),
			},
		},
		{
			name: "whitespace id",
			u: impl.NewUser(
				user.NewID("   "),
			),
			want: want{
				value: user.NewID("   "),
			},
		},
		{
			name: "default value",
			u:    impl.User{},
			want: want{
				value: user.NewID(""),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.u.ID(),
			)
		})
	}
}

func TestUser_Equals(t *testing.T) {
	t.Parallel()
	type args struct {
		other impl.User
	}
	type want struct {
		ok bool
	}
	tests := []struct {
		name string
		u    impl.User
		args args
		want want
	}{
		{
			name: "same with id",
			u: impl.NewUser(
				user.NewID("65e9c0b1a2b3c4d5e6f7a8b9"),
			),
			args: args{
				other: impl.NewUser(
					user.NewID("65e9c0b1a2b3c4d5e6f7a8b9"),
				),
			},
			want: want{
				ok: true,
			},
		},
		{
			name: "same with empty id",
			u: impl.NewUser(
				user.NewID(""),
			),
			args: args{
				other: impl.NewUser(
					user.NewID(""),
				),
			},
			want: want{
				ok: true,
			},
		},
		{
			name: "diff with id",
			u: impl.NewUser(
				user.NewID("65e9c0b1a2b3c4d5e6f7a8b9"),
			),
			args: args{
				other: impl.NewUser(
					user.NewID("65e9c0b1a2b3c4d5e6f7a8b0"),
				),
			},
			want: want{
				ok: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.u.Equals(tt.args.other)
			if tt.want.ok {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}
