package body_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/users/one/resp/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
)

func TestBody_User(t *testing.T) {
	t.Parallel()
	type want struct {
		val user.User
	}
	tests := []struct {
		name string
		b    body.Body
		want want
	}{
		{
			name: "ok",
			b: body.NewBody(
				userfx.NewAlexSmith(),
			),
			want: want{
				val: userfx.NewAlexSmith(),
			},
		},
		{
			name: "default value",
			b:    body.Body{},
			want: want{
				val: user.User{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.b.User(),
			)
		})
	}
}
