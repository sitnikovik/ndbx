package log_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
)

func TestDuration(t *testing.T) {
	t.Parallel()
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1 second",
			args: args{
				d: time.Second,
			},
			want: "\033[33m\"1s\"\033[0m",
		},
		{
			name: "1 minute",
			args: args{
				d: time.Minute,
			},
			want: "\033[33m\"1m0s\"\033[0m",
		},
		{
			name: "1 hour",
			args: args{
				d: time.Hour,
			},
			want: "\033[33m\"1h0m0s\"\033[0m",
		},
		{
			name: "1 day",
			args: args{
				d: 24 * time.Hour,
			},
			want: "\033[33m\"24h0m0s\"\033[0m",
		},
		{
			name: "1.5 hours",
			args: args{
				d: time.Hour + 30*time.Minute,
			},
			want: "\033[33m\"1h30m0s\"\033[0m",
		},
		{
			name: "0 sec",
			args: args{
				d: 0,
			},
			want: "\033[33m\"0s\"\033[0m",
		},
		{
			name: "negative duration",
			args: args{
				d: -time.Minute,
			},
			want: "\033[33m\"-1m0s\"\033[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want,
				log.Duration(tt.args.d),
			)
		})
	}
}
