package include_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint/rq/include"
)

func TestInclude_URLQuery(t *testing.T) {
	t.Parallel()
	type want struct {
		value url.Values
	}
	tests := []struct {
		name string
		i    impl.Include
		want want
	}{
		{
			name: "one inc",
			i:    impl.NewInclude("foo"),
			want: want{
				value: func() url.Values {
					v := make(url.Values, 1)
					v.Set("include", "foo")
					return v
				}(),
			},
		},
		{
			name: "two incs",
			i:    impl.NewInclude("foo", "bar"),
			want: want{
				value: func() url.Values {
					v := make(url.Values, 1)
					v.Set("include", "foo,bar")
					return v
				}(),
			},
		},
		{
			name: "empty",
			i:    impl.NewInclude(),
			want: want{
				value: func() url.Values {
					v := make(url.Values, 1)
					return v
				}(),
			},
		},
		{
			name: "space",
			i:    impl.NewInclude(" "),
			want: want{
				value: func() url.Values {
					v := make(url.Values, 1)
					v.Set("include", " ")
					return v
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				tt.i.URLQuery(),
			)
		})
	}
}
