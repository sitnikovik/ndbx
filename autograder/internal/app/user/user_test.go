package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

func TestUser_Hash(t *testing.T) {
	t.Parallel()
	type want struct {
		val string
	}
	tests := []struct {
		name string
		u    user.User
		want want
	}{
		{
			name: "ok",
			u: user.NewUser(
				user.NewID("1"),
				"samsep1ol",
				"Sam Sepiol",
			),
			want: want{
				val: "ad94cb9f636b4de5b18ee8526ac0d21d",
			},
		},
		{
			name: "default value",
			u:    user.User{},
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
				tt.u.Hash(),
			)
		})
	}
}
