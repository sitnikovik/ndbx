package body_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/pagination"
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
			name: "all fields",
			b: body.NewBody(
				body.WithTitle("test"),
				body.WithPagination(
					pagination.NewPagination(10, 10),
				),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 3)
					q.Set("title", "test")
					q.Set("limit", "10")
					q.Set("offset", "10")
					return q
				}(),
			},
		},
		{
			name: "only title",
			b: body.NewBody(
				body.WithTitle("test"),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 1)
					q.Set("title", "test")
					return q
				}(),
			},
		},
		{
			name: "only pagination",
			b: body.NewBody(
				body.WithPagination(
					pagination.NewPagination(10, 10),
				),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 2)
					q.Set("limit", "10")
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
