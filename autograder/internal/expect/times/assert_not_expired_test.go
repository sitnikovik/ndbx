package times_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/expect/times"
)

func TestAssertNotExpired(t *testing.T) {
	t.Parallel()
	type args struct {
		since    time.Time
		expected time.Duration
		actual   time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "not expired",
			args: args{
				since:    time.Now().Add(-time.Minute),
				expected: 2 * time.Minute,
				actual:   time.Second,
			},
			wantErr: nil,
		},
		{
			name: "expired",
			args: args{
				since:    time.Now().Add(-3 * time.Minute),
				expected: 2 * time.Minute,
				actual:   time.Second,
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "exactly expired",
			args: args{
				since:    time.Now().Add(-3 * time.Minute),
				expected: 2 * time.Minute,
				actual:   time.Minute,
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "just not expired",
			args: args{
				since:    time.Now().Add(-3*time.Minute + time.Second),
				expected: 2 * time.Minute,
				actual:   time.Minute,
			},
			wantErr: nil,
		},
		{
			name: "just expired",
			args: args{
				since:    time.Now().Add(-3*time.Minute - time.Second),
				expected: 2 * time.Minute,
				actual:   time.Minute,
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "zero expected and actual",
			args: args{
				since:    time.Now().Add(-time.Second),
				expected: 0,
				actual:   0,
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "negative expected and actual",
			args: args{
				since:    time.Now().Add(-time.Second),
				expected: -time.Minute,
				actual:   -time.Second,
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "negative expected and positive actual",
			args: args{
				since:    time.Now().Add(-time.Second),
				expected: -time.Minute,
				actual:   time.Second,
			},
			wantErr: errs.ErrExpectationFailed,
		},
		{
			name: "positive expected and negative actual",
			args: args{
				since:    time.Now().Add(-time.Second),
				expected: time.Minute,
				actual:   -time.Second,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := times.AssertNotExpired(
				tt.args.since,
				tt.args.expected,
				tt.args.actual,
			)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
