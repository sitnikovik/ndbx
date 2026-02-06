package log_test

import (
	"testing"

	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	t.Parallel()
	type args struct {
		u string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok url",
			args: args{
				u: "https://example.com",
			},
			want: "\033[32m\"https://example.com\"\033[0m",
		},
		{
			name: "empty url",
			args: args{
				u: "",
			},
			want: "\033[32m\"\"\033[0m",
		},
		{
			name: "url with spaces",
			args: args{
				u: " \"https://example.com\" ",
			},
			want: "\033[32m\" \"https://example.com\" \"\033[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := log.URL(tt.args.u)
			assert.Equal(t, tt.want, got)
		})
	}
}
