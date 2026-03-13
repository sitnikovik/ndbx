package strings_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/strings"
)

func TestAssertNotEquals(t *testing.T) {
	t.Parallel()
	type args struct {
		expected string
		actual   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "ok",
			args: args{
				expected: "hello",
				actual:   "hello",
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "not equal",
			args: args{
				expected: "hello",
				actual:   "world",
			},
			wantErr: nil,
		},
		{
			name: "empty strings",
			args: args{
				expected: "",
				actual:   "",
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "expected empty but actual not empty",
			args: args{
				expected: "",
				actual:   "not empty",
			},
			wantErr: nil,
		},
		{
			name: "expected not empty but actual empty",
			args: args{
				expected: "not empty",
				actual:   "",
			},
			wantErr: nil,
		},
		{
			name: "case sensitive",
			args: args{
				expected: "Hello",
				actual:   "hello",
			},
			wantErr: nil,
		},
		{
			name: "whitespace matters",
			args: args{
				expected: "hello ",
				actual:   "hello",
			},
			wantErr: nil,
		},
		{
			name: "multiline strings",
			args: args{
				expected: "line1\nline2",
				actual:   "line1\nline2",
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "multiline strings not equal",
			args: args{
				expected: "line1\nline2",
				actual:   "line1\nline3",
			},
			wantErr: nil,
		},
		{
			name: "strings with special characters",
			args: args{
				expected: "hello\tworld",
				actual:   "hello\tworld",
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "strings with special characters not equal",
			args: args{
				expected: "hello\tworld",
				actual:   "hello world",
			},
			wantErr: nil,
		},
		{
			name: "unicode strings equal",
			args: args{
				expected: "привет",
				actual:   "привет",
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "unicode strings not equal",
			args: args{
				expected: "привет",
				actual:   "пока",
			},
			wantErr: nil,
		},
		{
			name: "all empty",
			args: args{
				expected: "",
				actual:   "",
			},
			wantErr: errs.ErrExpectationFailed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.ErrorIs(
				t,
				strings.AssertNotEquals(
					tt.args.expected,
					tt.args.actual,
				),
				tt.wantErr,
			)
		})
	}
}
