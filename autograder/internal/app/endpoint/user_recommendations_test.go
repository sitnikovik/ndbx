package endpoint_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/sitnikovik/ndbx/autograder/internal/app/endpoint"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
)

func TestUserRecommendations(t *testing.T) {
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
				value: "/users/123/recommendations",
			},
		},
		{
			name: "empty user id",
			args: args{
				id: user.NewID(""),
			},
			want: want{
				value: "/users//recommendations",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want.value,
				impl.UserRecommendations(tt.args.id),
			)
		})
	}
}
