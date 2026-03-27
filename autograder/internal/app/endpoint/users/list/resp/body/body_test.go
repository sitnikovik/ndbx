package body_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/users/list/resp/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
)

func TestBody_Users(t *testing.T) {
	t.Parallel()
	type want struct {
		val []user.User
	}
	tests := []struct {
		name string
		b    body.Body
		want want
	}{
		{
			name: "ok",
			b: body.NewBody(
				[]user.User{
					userfx.NewJohnDoe(),
				},
				1,
			),
			want: want{
				val: []user.User{
					userfx.NewJohnDoe(),
				},
			},
		},
		{
			name: "empty",
			b: body.NewBody(
				[]user.User{},
				0,
			),
			want: want{
				val: []user.User{},
			},
		},
		{
			name: "has users but count is zero",
			b: body.NewBody(
				[]user.User{
					userfx.NewJohnDoe(),
				},
				0,
			),
			want: want{
				val: []user.User{
					userfx.NewJohnDoe(),
				},
			},
		},
		{
			name: "count is gt zero but has no users",
			b: body.NewBody(
				[]user.User{},
				1,
			),
			want: want{
				val: []user.User{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.b.Users(),
			)
		})
	}
}

func TestBody_Count(t *testing.T) {
	t.Parallel()
	type want struct {
		val int
	}
	tests := []struct {
		name string
		b    body.Body
		want want
	}{
		{
			name: "ok",
			b: body.NewBody(
				[]user.User{
					userfx.NewJohnDoe(),
				},
				1,
			),
			want: want{
				val: 1,
			},
		},
		{
			name: "empty",
			b: body.NewBody(
				[]user.User{},
				0,
			),
			want: want{
				val: 0,
			},
		},
		{
			name: "has users but count is zero",
			b: body.NewBody(
				[]user.User{
					userfx.NewJohnDoe(),
				},
				0,
			),
			want: want{
				val: 0,
			},
		},
		{
			name: "count is gt zero but has no users",
			b: body.NewBody(
				[]user.User{},
				1,
			),
			want: want{
				val: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.b.Count(),
			)
		})
	}
}
