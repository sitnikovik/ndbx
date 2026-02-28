package redis_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/health/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	redisfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/redis"
)

func TestStep_Run(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx  context.Context
		vars step.Variables
	}
	tests := []struct {
		name      string
		step      *redis.Step
		args      args
		wantErr   error
		wantVars  step.Variables
		wantPanic bool
	}{
		{
			name: "ok",
			step: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithEmpty(
						func(_ context.Context) (bool, error) {
							return true, nil
						},
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			wantErr:   nil,
			wantVars:  step.NewVariables(),
			wantPanic: false,
		},
		{
			name: "redis failed",
			step: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithEmpty(
						func(_ context.Context) (bool, error) {
							return false, assert.AnError
						},
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			wantErr:   errs.ErrRedisFailed,
			wantVars:  step.NewVariables(),
			wantPanic: false,
		},
		{
			name: "redis is not empty",
			step: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithEmpty(
						func(_ context.Context) (bool, error) {
							return false, nil
						},
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			wantErr:   errs.ErrExpectationFailed,
			wantVars:  step.NewVariables(),
			wantPanic: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.wantPanic {
				assert.Panics(
					t,
					func() {
						_ = tt.step.Run(
							tt.args.ctx,
							tt.args.vars,
						)
					},
				)
			} else {
				assert.ErrorIs(
					t,
					tt.step.Run(
						tt.args.ctx,
						tt.args.vars,
					),
					tt.wantErr,
				)
			}
			assert.Equal(
				t,
				tt.wantVars,
				tt.args.vars,
			)
		})
	}
}
