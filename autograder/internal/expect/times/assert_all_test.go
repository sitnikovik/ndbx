package times_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/times"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestAssertAll(t *testing.T) {
	t.Parallel()
	type args struct {
		expected time.Time
		actual   time.Time
		ff       []times.AssertFunc
	}
	type want struct {
		err   error
		panic bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{

		{
			name: "all assertions pass",
			args: args{
				expected: timex.MustParse(time.RFC3339, "2024-01-01T00:00:00Z"),
				actual:   timex.MustParse(time.RFC3339, "2024-01-01T00:00:00Z"),
				ff: []times.AssertFunc{
					times.AssertEquals,
					times.AssertNotAfter,
				},
			},
			want: want{
				err:   nil,
				panic: false,
			},
		},
		{
			name: "one assertion fails",
			args: args{
				expected: timex.MustParse(time.RFC3339, "2024-01-01T00:00:00Z"),
				actual:   timex.MustParse(time.RFC3339, "2024-01-02T00:00:00Z"),
				ff: []times.AssertFunc{
					times.AssertEquals,
					times.AssertNotAfter,
				},
			},
			want: want{
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "no assertion functions provided",
			args: args{
				expected: timex.MustParse(time.RFC3339, "2024-01-01T00:00:00Z"),
				actual:   timex.MustParse(time.RFC3339, "2024-01-01T00:00:00Z"),
				ff:       []times.AssertFunc{},
			},
			want: want{
				err:   nil,
				panic: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want.panic {
				assert.Panics(
					t,
					func() {
						_ = times.AssertAll(
							tt.args.expected,
							tt.args.actual,
							tt.args.ff...,
						)
					},
				)
				return
			}
			assert.ErrorIs(
				t,
				times.AssertAll(
					tt.args.expected,
					tt.args.actual,
					tt.args.ff...,
				),
				tt.want.err,
			)
		})
	}
}
