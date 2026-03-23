package rangeof_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	rangeof "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/range-of"
)

func TestInts_URLQuery(t *testing.T) {
	t.Parallel()
	type want struct {
		val url.Values
	}
	tests := []struct {
		name string
		r    rangeof.Ints
		want want
	}{
		{
			name: "all are set",
			r: rangeof.NewInts(
				"price",
				rangeof.NewInt(-100),
				rangeof.NewInt(300),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 2)
					q.Set("price_from", "-100")
					q.Set("price_to", "300")
					return q
				}(),
			},
		},
		{
			name: "only from is set",
			r: rangeof.NewInts(
				"price",
				rangeof.NewInt(100),
				rangeof.NewInt(0),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 1)
					q.Set("price_from", "100")
					return q
				}(),
			},
		},
		{
			name: "only to is set",
			r: rangeof.NewInts(
				"price",
				rangeof.NewInt(0),
				rangeof.NewInt(-300),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 1)
					q.Set("price_to", "-300")
					return q
				}(),
			},
		},
		{
			name: "all are empty",
			r: rangeof.NewInts(
				"price",
				rangeof.NewInt(0),
				rangeof.NewInt(0),
			),
			want: want{
				val: func() url.Values {
					return make(url.Values)
				}(),
			},
		},
		{
			name: "all are set as any",
			r: rangeof.NewInts(
				"price",
				rangeof.NewAnyInt(0),
				rangeof.NewAnyInt(0),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 2)
					q.Set("price_from", "0")
					q.Set("price_to", "0")
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
				tt.r.URLQuery(),
			)
		})
	}
}
