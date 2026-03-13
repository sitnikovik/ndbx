package times_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/times"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestAssertAfterOrEqual(t *testing.T) {
	t.Parallel()
	type args struct {
		expected time.Time
		actual   time.Time
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
			name: "actual is after expected",
			args: args{
				expected: timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
				actual:   timex.MustParse(time.RFC3339, "2024-06-02T00:00:00Z"),
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "actual is equal to expected",
			args: args{
				expected: timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
				actual:   timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "actual is before expected",
			args: args{
				expected: timex.MustParse(time.RFC3339, "2024-06-02T00:00:00Z"),
				actual:   timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
			},
			want: want{
				err: errs.ErrExpectationFailed,
			},
		},
		{
			name: "actual is before expected with different timezones",
			args: args{
				expected: timex.MustParse(time.RFC3339, "2024-06-01T00:00:00+02:00"),
				actual:   timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "actual is after expected with different timezones",
			args: args{
				expected: timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
				actual:   timex.MustParse(time.RFC3339, "2024-06-01T00:00:00+02:00"),
			},
			want: want{
				err: errs.ErrExpectationFailed,
			},
		},
		{
			name: "actual is equal to expected with different timezones",
			args: args{
				expected: timex.MustParse(time.RFC3339, "2024-06-01T00:00:00+02:00"),
				actual:   timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "actual is equal to expected with different timezones but same time",
			args: args{
				expected: timex.MustParse(time.RFC3339, "2024-06-01T00:00:00+02:00"),
				actual:   timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "actual is before expected with different timezones but same time",
			args: args{
				expected: timex.MustParse(time.RFC3339, "2024-06-01T00:00:00Z"),
				actual:   timex.MustParse(time.RFC3339, "2024-06-01T00:00:00+02:00"),
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
				times.AssertAfterOrEqual(
					tt.args.expected,
					tt.args.actual,
				),
				tt.want.err,
			)
		})
	}
}
