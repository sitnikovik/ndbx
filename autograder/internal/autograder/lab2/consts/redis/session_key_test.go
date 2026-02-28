package redis_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/redis"
)

func TestSessionKey(t *testing.T) {
	t.Parallel()
	type args struct {
		sessionID string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple session ID",
			args: args{
				sessionID: "abc123",
			},
			want: "sid:abc123",
		},
		{
			name: "empty session ID",
			args: args{
				sessionID: "",
			},
			want: "sid:",
		},
		{
			name: "session ID with special characters",
			args: args{
				sessionID: "user!@#$%^&*()",
			},
			want: "sid:user!@#$%^&*()",
		},
		{
			name: "session ID with spaces",
			args: args{
				sessionID: "session id with spaces",
			},
			want: "sid:session id with spaces",
		},
		{
			name: "empty session id",
			args: args{
				sessionID: "",
			},
			want: "sid:",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := redis.SessionKey(tt.args.sessionID)
			assert.Equal(t, tt.want, got)
		})
	}
}
