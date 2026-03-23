package rangeof_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	rangeof "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/range-of"
)

func TestStrings_URLQuery(t *testing.T) {
	t.Parallel()
	type want struct {
		val url.Values
	}
	tests := []struct {
		name string
		r    rangeof.Strings
		want want
	}{
		{
			name: "all are set",
			r: rangeof.NewStrings(
				"label",
				"foo",
				"bar",
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 2)
					q.Set("label_from", "foo")
					q.Set("label_to", "bar")
					return q
				}(),
			},
		},
		{
			name: "only from is set",
			r: rangeof.NewStrings(
				"label",
				"foo",
				"",
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 1)
					q.Set("label_from", "foo")
					return q
				}(),
			},
		},
		{
			name: "only to is set",
			r: rangeof.NewStrings(
				"label",
				"",
				"bar",
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 1)
					q.Set("label_to", "bar")
					return q
				}(),
			},
		},
		{
			name: "all are empty",
			r: rangeof.NewStrings(
				"label",
				"",
				"",
			),
			want: want{
				val: func() url.Values {
					return make(url.Values)
				}(),
			},
		},
		{
			name: "spaces",
			r: rangeof.NewStrings(
				"label",
				" ",
				" ",
			),
			want: want{
				val: func() url.Values {
					return make(url.Values)
				}(),
			},
		},
		{
			name: "unicode",
			r: rangeof.NewStrings(
				"label",
				"привет",
				"мир",
			),
			want: want{
				val: func() url.Values {
					q := make(url.Values, 2)
					q.Set("label_from", "привет")
					q.Set("label_to", "мир")
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
