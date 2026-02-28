package duration_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/times/duration"
)

func TestAssertEquals(t *testing.T) {
	t.Parallel()
	type args struct {
		expected time.Duration
		actual   time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "equal durations",
			args: args{
				expected: time.Minute,
				actual:   time.Minute,
			},
			wantErr: nil,
		},
		{
			name: "different durations",
			args: args{
				expected: time.Minute,
				actual:   2 * time.Minute,
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "zero durations",
			args: args{
				expected: 0,
				actual:   0,
			},
			wantErr: nil,
		},
		{
			name: "negative durations",
			args: args{
				expected: -time.Minute,
				actual:   -time.Minute,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.ErrorIs(
				t,
				duration.AssertEquals(
					tt.args.expected,
					tt.args.actual,
				),
				tt.wantErr,
			)
		})
	}
}
