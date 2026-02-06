package log_test

import (
	"testing"

	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok string",
			args: args{
				s: "test",
			},
			want: "\033[33m\"test\"\033[0m",
		},
		{
			name: "empty string",
			args: args{
				s: "",
			},
			want: "\033[33m\"\"\033[0m",
		},
		{
			name: "string with spaces",
			args: args{
				s: " test ",
			},
			want: "\033[33m\" test \"\033[0m",
		},
		{
			name: "only space",
			args: args{
				s: " ",
			},
			want: "\033[33m\" \"\033[0m",
		},
		{
			name: "string with special characters",
			args: args{
				s: "test\nwith\tspecial\"characters\"",
			},
			want: "\033[33m\"test\nwith\tspecial\"characters\"\"\033[0m",
		},
		{
			name: "string with unicode characters",
			args: args{
				s: "тест",
			},
			want: "\033[33m\"тест\"\033[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := log.String(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
