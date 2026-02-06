package log_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/paints/log"
)

func TestBool(t *testing.T) {
	t.Parallel()
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "true",
			args: args{b: true},
			want: "\033[35mtrue\033[0m",
		},
		{
			name: "false",
			args: args{
				b: false,
			},
			want: "\033[35mfalse\033[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := log.Bool(tt.args.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
