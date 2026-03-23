package rangeof_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	rangeof "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/range-of"
)

func TestUInts_URLQuery(t *testing.T) {
	t.Parallel()
	type want struct {
		val url.Values
	}
	tests := []struct {
		name string
		r    rangeof.UInts
		want want
	}{
		{
			name: "all are set",
			r: rangeof.NewUInts(
				"age",
				rangeof.NewUInt(uint(18)),
				rangeof.NewUInt(uint(30)),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 2)
					q.Set("age_from", "18")
					q.Set("age_to", "30")
					return q
				}(),
			},
		},
		{
			name: "only from is set",
			r: rangeof.NewUInts(
				"age",
				rangeof.NewUInt(uint(18)),
				rangeof.NewUInt(uint(30)),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 1)
					q.Set("age_from", "18")
					return q
				}(),
			},
		},
		{
			name: "only to is set",
			r: rangeof.NewUInts(
				"age",
				rangeof.NewUInt(uint(0)),
				rangeof.NewUInt(uint(30)),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 1)
					q.Set("age_to", "30")
					return q
				}(),
			},
		},
		{
			name: "all are empty",
			r: rangeof.NewUInts(
				"age",
				rangeof.NewUInt(uint(0)),
				rangeof.NewUInt(uint(0)),
			),
			want: want{
				val: func() url.Values {
					return make(url.Values)
				}(),
			},
		},
		{
			name: "all are set as any",
			r: rangeof.NewUInts(
				"age",
				rangeof.NewAnyUInt(uint(0)),
				rangeof.NewAnyUInt(uint(0)),
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 2)
					q.Set("age_from", "0")
					q.Set("age_to", "0")
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
