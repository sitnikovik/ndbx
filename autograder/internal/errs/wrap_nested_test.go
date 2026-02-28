package errs_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/errs"
)

func TestWrapNested(t *testing.T) {
	t.Parallel()
	type args struct {
		curr   error
		prev   error
		format string
		args   []any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				curr:   context.DeadlineExceeded,
				prev:   assert.AnError,
				format: "additional context: %s",
				args:   []any{"details"},
			},
			wantErr: true,
		},
		{
			name: "nil current error",
			args: args{
				curr:   nil,
				prev:   assert.AnError,
				format: "additional context: %s",
				args:   []any{"details"},
			},
			wantErr: false,
		},
		{
			name: "nil previous error",
			args: args{
				curr:   assert.AnError,
				prev:   nil,
				format: "additional context: %s",
				args:   []any{"details"},
			},
			wantErr: false,
		},
		{
			name: "both errors nil",
			args: args{
				curr:   nil,
				prev:   nil,
				format: "additional context: %s",
				args:   []any{"details"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := errs.WrapNested(
				tt.args.curr,
				tt.args.prev,
				tt.args.format,
				tt.args.args...,
			)
			if tt.wantErr {
				assert.ErrorIs(t, err, tt.args.curr)
				assert.NotErrorIs(t, err, tt.args.prev)
				assert.Contains(t, err.Error(), fmt.Sprintf(tt.args.format, tt.args.args...))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
