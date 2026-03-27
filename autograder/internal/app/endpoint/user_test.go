package endpoint_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
)

func TestEndpoint_User(t *testing.T) {
	t.Parallel()
	type args struct {
		id string
	}
	type want struct {
		val string
	}
	tests := []struct {
		name string
		e    endpoint.Endpoint
		args args
		want want
	}{
		{
			name: "ok",
			e:    endpoint.NewEndpoint("http://localhost"),
			args: args{
				id: "123",
			},
			want: want{
				val: "http://localhost/users/123",
			},
		},
		{
			name: "empty id",
			e:    endpoint.NewEndpoint("http://localhost"),
			args: args{
				id: "",
			},
			want: want{
				val: "http://localhost/users/",
			},
		},
		{
			name: "empty base URL",
			e:    endpoint.NewEndpoint(""),
			args: args{
				id: "123",
			},
			want: want{
				val: "/users/123",
			},
		},
		{
			name: "empty base URL and id",
			e:    endpoint.NewEndpoint(""),
			args: args{
				id: "",
			},
			want: want{
				val: "/users/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.val,
				tt.e.User(tt.args.id),
			)
		})
	}
}
