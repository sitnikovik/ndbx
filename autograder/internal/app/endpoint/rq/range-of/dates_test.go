package rangeof_test

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	rangeof "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/range-of"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestDates_URLQuery(t *testing.T) {
	t.Parallel()
	type want struct {
		val url.Values
	}
	tests := []struct {
		name string
		r    rangeof.Dates
		want want
	}{
		{
			name: "all are set",
			r: rangeof.NewDates(
				"date",
				timex.MustRFC3339("2025-01-01T11:00:00Z"),
				timex.MustRFC3339("2025-01-07T11:00:00Z"),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 2)
					q.Set("date_from", "20250101")
					q.Set("date_to", "20250107")
					return q
				}(),
			},
		},
		{
			name: "only from is set",
			r: rangeof.NewDates(
				"date",
				timex.MustRFC3339("2025-01-01T11:00:00Z"),
				time.Time{},
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 1)
					q.Set("date_from", "20250101")
					return q
				}(),
			},
		},
		{
			name: "only to is set",
			r: rangeof.NewDates(
				"date",
				time.Time{},
				timex.MustRFC3339("2025-01-07T11:00:00Z"),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 1)
					q.Set("date_to", "20250107")
					return q
				}(),
			},
		},
		{
			name: "all are empty",
			r: rangeof.NewDates(
				"date",
				time.Time{},
				time.Time{},
			),
			want: want{
				val: func() url.Values {
					return make(url.Values)
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
				tt.r.URLQuery(),
			)
		})
	}
}
