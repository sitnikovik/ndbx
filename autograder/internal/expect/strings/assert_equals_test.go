package strings_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/strings"
)

func TestAssertEquals(t *testing.T) {
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
			wantErr: nil,
		},
		{
			name: "not equal",
			args: args{
				expected: "hello",
				actual:   "world",
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "empty strings",
			args: args{
				expected: "",
				actual:   "",
			},
			wantErr: nil,
		},
		{
			name: "expected empty but actual not empty",
			args: args{
				expected: "",
				actual:   "not empty",
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "expected not empty but actual empty",
			args: args{
				expected: "not empty",
				actual:   "",
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "case sensitive",
			args: args{
				expected: "Hello",
				actual:   "hello",
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "whitespace matters",
			args: args{
				expected: "hello ",
				actual:   "hello",
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "multiline strings",
			args: args{
				expected: "line1\nline2",
				actual:   "line1\nline2",
			},
			wantErr: nil,
		},
		{
			name: "multiline strings not equal",
			args: args{
				expected: "line1\nline2",
				actual:   "line1\nline3",
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "strings with special characters",
			args: args{
				expected: "hello\tworld",
				actual:   "hello\tworld",
			},
			wantErr: nil,
		},
		{
			name: "strings with special characters not equal",
			args: args{
				expected: "hello\tworld",
				actual:   "hello world",
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "unicode strings equal",
			args: args{
				expected: "привет",
				actual:   "привет",
			},
			wantErr: nil,
		},
		{
			name: "unicode strings not equal",
			args: args{
				expected: "привет",
				actual:   "пока",
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "all empty",
			args: args{
				expected: "",
				actual:   "",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.ErrorIs(
				t,
				strings.AssertEquals(
					tt.args.expected,
					tt.args.actual,
				),
				tt.wantErr,
			)
		})
	}
}
