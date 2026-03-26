package expect_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect"
)

func TestAssertEquals(t *testing.T) {
	t.Parallel()
	type args struct {
		expect any
		actual any
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
			name: "same ints",
			args: args{
				expect: 1,
				actual: 1,
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "diff ints",
			args: args{
				expect: 1,
				actual: 0,
			},
			want: want{
				err: errs.ErrExpectationFailed,
			},
		},
		{
			name: "same strings",
			args: args{
				expect: "foo",
				actual: "foo",
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "diff strings",
			args: args{
				expect: "foo",
				actual: "bar",
			},
			want: want{
				err: errs.ErrExpectationFailed,
			},
		},
		{
			name: "same slices",
			args: args{
				expect: []string{"1"},
				actual: []string{"1"},
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "diff slices",
			args: args{
				expect: []string{"1"},
				actual: []string{"1", "1"},
			},
			want: want{
				err: errs.ErrExpectationFailed,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.ErrorIs(
				t,
				expect.AssertEquals(
					tt.args.expect,
					tt.args.actual,
				),
				tt.want.err,
			)
		})
	}
}
