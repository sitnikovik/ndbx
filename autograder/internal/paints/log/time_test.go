package log_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestTime(t *testing.T) {
	t.Parallel()
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "date time",
			args: args{
				t: timex.MustParse(time.DateTime, "2024-06-01 12:00:00"),
			},
			want: "\033[33m\"2024-06-01T12:00:00Z\"\033[0m",
		},
		{
			name: "rfc3339",
			args: args{
				t: timex.MustParse(time.RFC3339, "2024-06-01T12:00:00+03:00"),
			},
			want: "\033[33m\"2024-06-01T12:00:00+03:00\"\033[0m",
		},
		{
			name: "unix utc time",
			args: args{
				t: time.Unix(1700000000, 0).UTC(),
			},
			want: "\033[33m\"2023-11-14T22:13:20Z\"\033[0m",
		},
		{
			name: "zero time",
			args: args{
				t: time.Time{},
			},
			want: "\033[33m\"0001-01-01T00:00:00Z\"\033[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want,
				log.Time(tt.args.t),
			)
		})
	}
}
