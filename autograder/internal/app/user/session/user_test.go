package session_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user/session"
)

func TestUser_ID(t *testing.T) {
	t.Parallel()
	type want struct {
		val user.ID
	}
	tests := []struct {
		name string
		u    session.User
		want want
	}{
		{
			name: "ok",
			u: session.NewUser(
				user.NewID("1"),
			),
			want: want{
				val: user.NewID("1"),
			},
		},
		{
			name: "empty ID",
			u: session.NewUser(
				user.NewID(""),
			),
			want: want{
				val: user.NewID(""),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.u.ID(),
			)
		})
	}
}
