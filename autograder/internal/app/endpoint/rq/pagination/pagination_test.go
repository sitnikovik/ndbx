package pagination_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/pagination"
)

func TestPagination_URLQuery(t *testing.T) {
	t.Parallel()
	type want struct {
		val url.Values
	}
	tests := []struct {
		name string
		p    pagination.Pagination
		want want
	}{
		{
			name: "ok",
			p:    pagination.NewPagination(10, 5),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 2)
					q.Set("limit", "10")
					q.Set("offset", "5")
					return q
				}(),
			},
		},
		{
			name: "zero values",
			p:    pagination.NewPagination(0, 0),
			want: want{
				val: make(url.Values),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.p.URLQuery(),
			)
		})
	}
}
