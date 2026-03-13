package strings_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/strings"
)

func TestAssertNotEmpty(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	type want struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "non-empty string",
			args: args{
				s: "hello",
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "empty string",
			args: args{
				s: "",
			},
			want: want{
				err: errs.ErrExpectationFailed,
			},
		},
		{
			name: "string with spaces",
			args: args{
				s: "   ",
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "string with newline",
			args: args{
				s: "\n",
			},
			want: want{
				err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.ErrorIs(
				t,
				strings.AssertNotEmpty(
					tt.args.s,
				),
				tt.want.err,
			)
		})
	}
}
