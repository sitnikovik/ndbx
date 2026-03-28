package body_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/events/get/rq/body"
	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/pagination"
	"github.com/sitnikovik/ndbx/autograder/internal/app/event/category"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
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
				body.WithCategory(category.Concert),
				body.WithAddress("NY, Groove street, 123/1"),
				body.WithCity("New York"),
				body.WithEntryPrice(0, 0),
				body.WithDates(
					timex.MustRFC3339("2025-03-01T11:00:00Z"),
					timex.MustRFC3339("2025-03-07T11:00:00Z"),
				),
				body.WithByUser(
					user.NewIdentity(
						user.NewID("123"),
						user.WithUsername("username"),
					),
				),
				body.WithPagination(
					pagination.NewPagination(10, 10),
				),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 11)
					q.Set("title", "test")
					q.Set("category", category.Concert.String())
					q.Set("address", "NY, Groove street, 123/1")
					q.Set("city", "New York")
					q.Set("price_to", "0")
					q.Set("date_from", "20250301")
					q.Set("date_to", "20250307")
					q.Set("user_id", "123")
					q.Set("user", "username")
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
