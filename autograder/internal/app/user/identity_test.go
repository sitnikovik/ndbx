package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

func TestIdentity_Username(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		i    user.Identity
		want want
	}{
		{
			name: "ok",
			i: user.NewIdentity(
				user.NewID("1"),
				user.WithUsername("username"),
			),
			want: want{
				val: "username",
			},
		},
		{
			name: "only id",
			i: user.NewIdentity(
				user.NewID("1"),
			),
			want: want{
				val: "",
			},
		},
		{
			name: "default value",
			i:    user.Identity{},
			want: want{
				val: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.i.Username(),
			)
		})
	}
}

func TestIdentity_Hash(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		i    user.Identity
		want want
	}{
		{
			name: "ok",
			i: user.NewIdentity(
				user.NewID("1"),
				user.WithUsername(
					"samsep1ol",
				),
			),
			want: want{
				val: "ad94cb9f636b4de5b18ee8526ac0d21d",
			},
		},
		{
			name: "default value",
			i:    user.Identity{},
			want: want{
				val: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.i.Hash(),
			)
		})
	}
}
