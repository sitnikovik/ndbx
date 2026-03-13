package redis_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/auth/logout/ok/redis"
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
					redisfk.WithHas(
						func(
							_ context.Context,
							_ string,
						) (bool, error) {
							return false, nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						session.Name,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
			},
			wantErr: nil,
			wantVars: func() step.Variables {
				vars := step.NewVariables()
				vars.Set(
					session.Name,
					"0123456789abcdef0123456789abcdef",
				)
				return vars
			}(),
			wantPanic: false,
		},
		{
			name: "missed session name variable",
			step: redis.NewStep(
				redisfk.NewFakeClient(),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			wantErr:   nil,
			wantVars:  step.NewVariables(),
			wantPanic: true,
		},
		{
			name: "redis failed",
			step: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHas(
						func(
							_ context.Context,
							_ string,
						) (bool, error) {
							return false, assert.AnError
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						session.Name,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
			},
			wantErr: assert.AnError,
			wantVars: func() step.Variables {
				vars := step.NewVariables()
				vars.Set(
					session.Name,
					"0123456789abcdef0123456789abcdef",
				)
				return vars
			}(),
			wantPanic: false,
		},
		{
			name: "session still exists",
			step: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHas(
						func(
							_ context.Context,
							_ string,
						) (bool, error) {
							return true, nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						session.Name,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
			},
			wantErr: errs.ErrExpectationFailed,
			wantVars: func() step.Variables {
				vars := step.NewVariables()
				vars.Set(
					session.Name,
					"0123456789abcdef0123456789abcdef",
				)
				return vars
			}(),
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
