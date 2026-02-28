package times_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/times"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

func TestAssertEquals(t *testing.T) {
	t.Parallel()
	type args struct {
		expected time.Time
		actual   time.Time
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "equal times",
			args: args{
				expected: timex.MustParse(
					time.RFC3339,
					"2024-06-01T12:00:00Z",
				),
				actual: timex.MustParse(
					time.RFC3339,
					"2024-06-01T12:00:00Z",
				),
			},
			wantErr: nil,
		},
		{
			name: "different times",
			args: args{
				expected: timex.MustParse(
					time.RFC3339,
					"2024-06-01T12:00:00Z",
				),
				actual: timex.MustParse(
					time.RFC3339,
					"2024-06-01T12:01:00Z",
				),
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "equal times with different time zones",
			args: args{
				expected: timex.MustParse(
					time.RFC3339,
					"2024-06-01T12:00:00+03:00",
				),
				actual: timex.MustParse(
					time.RFC3339,
					"2024-06-01T09:00:00Z",
				),
			},
			wantErr: nil,
		},
		{
			name: "different times with different time zones",
			args: args{
				expected: timex.MustParse(
					time.RFC3339,
					"2024-06-01T12:00:00+03:00",
				),
				actual: timex.MustParse(
					time.RFC3339,
					"2024-06-01T12:00:00Z",
				),
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "zero times",
			args: args{
				expected: time.Time{},
				actual:   time.Time{},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.ErrorIs(
				t,
				times.AssertEquals(
					tt.args.expected,
					tt.args.actual,
				),
				tt.wantErr,
			)
		})
	}
}
