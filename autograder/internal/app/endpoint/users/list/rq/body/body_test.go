package body_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/pagination"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/users/list/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

func TestBody_URLQuery(t *testing.T) {
	t.Parallel()
	type want struct {
		val url.Values
	}
	tests := []struct {
		name string
		b    body.Body
		want want
	}{
		{
			name: "all set",
			b: body.NewBody(
				body.WithIdentity(
					user.NewIdentity(
						user.ID("1"),
						user.WithUsername("john_doe"),
					),
				),
				body.WithFullName("John Doe"),
				body.WithPagination(
					pagination.NewPagination(
						10,
						20,
					),
				),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 4)
					q.Set("id", "1")
					q.Set("name", "John Doe")
					q.Set("limit", "10")
					q.Set("offset", "20")
					return q
				}(),
			},
		},
		{
			name: "only id",
			b: body.NewBody(
				body.WithIdentity(
					user.NewIdentity(
						user.ID("1"),
					),
				),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 1)
					q.Set("id", "1")
					return q
				}(),
			},
		},
		{
			name: "only username",
			b: body.NewBody(
				body.WithIdentity(
					user.NewIdentity(
						user.ID(""),
						user.WithUsername("john_doe"),
					),
				),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 0)
					return q
				}(),
			},
		},
		{
			name: "only name",
			b: body.NewBody(
				body.WithFullName("John Doe"),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 1)
					q.Set("name", "John Doe")
					return q
				}(),
			},
		},
		{
			name: "only limit",
			b: body.NewBody(
				body.WithPagination(
					pagination.NewPagination(
						10,
						0,
					),
				),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 1)
					q.Set("limit", "10")
					return q
				}(),
			},
		},
		{
			name: "only offset",
			b: body.NewBody(
				body.WithPagination(
					pagination.NewPagination(
						0,
						10,
					),
				),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 1)
					q.Set("offset", "10")
					return q
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.b.URLQuery(),
			)
		})
	}
}
