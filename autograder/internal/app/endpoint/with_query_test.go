package endpoint_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
)

func TestWithQuery(t *testing.T) {
	t.Parallel()
	type args struct {
		endpoint string
		q        url.Values
	}
	type want struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "ok",
			args: args{
				endpoint: "http://localhost/endpoint",
				q: func() url.Values {
					v := make(url.Values, 2)
					v.Set("key1", "value1")
					v.Set("key2", "value2")
					return v
				}(),
			},
			want: want{
				val: "http://localhost/endpoint?key1=value1&key2=value2",
			},
		},
		{
			name: "empty endpoint",
			args: args{
				endpoint: "",
				q: func() url.Values {
					v := make(url.Values, 1)
					v.Set("key", "value")
					return v
				}(),
			},
			want: want{
				val: "?key=value",
			},
		},
		{
			name: "empty query",
			args: args{
				endpoint: "http://localhost/endpoint",
				q:        make(url.Values),
			},
			want: want{
				val: "http://localhost/endpoint",
			},
		},
		{
			name: "without base url",
			args: args{
				endpoint: "/endpoint",
				q: func() url.Values {
					v := make(url.Values, 2)
					v.Set("key1", "value1")
					v.Set("key2", "value2")
					return v
				}(),
			},
			want: want{
				val: "/endpoint?key1=value1&key2=value2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				endpoint.WithQuery(
					tt.args.endpoint,
					tt.args.q,
				),
			)
		})
	}
}
