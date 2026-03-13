package lab3_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3"
)

func TestNewTestUser(t *testing.T) {
	t.Parallel()
	type want struct {
		val user.User
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "ok",
			want: want{
				val: user.NewUser(
					user.NewID(""),
					"sams3p1ol",
					"Sam Sepiol",
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				lab3.NewTestUser(),
			)
		})
	}
}
