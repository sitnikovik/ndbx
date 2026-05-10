package endpoint_test

import (
	"testing"

	"github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/stretchr/testify/assert"
)

func TestUserInterests(t *testing.T) {
	t.Parallel()
	type args struct {
		id user.ID
	}
	type want struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "ok",
			args: args{
				id: user.NewID("123"),
			},
			want: want{
				value: "/users/123/interests",
			},
		},
		{
			name: "empty id",
			args: args{
				id: user.NewID(""),
			},
			want: want{
				value: "/users//interests",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				endpoint.UserInterests(tt.args.id),
			)
		})
	}
}
