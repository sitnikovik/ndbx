package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/cookie"
	consts "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/session/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	redisfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
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
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								"created_at": "2026-01-01T00:00:00Z",
							}, nil
						},
					),
					redisfk.WithTTL(
						func(
							_ context.Context,
							_ string,
						) (time.Duration, error) {
							return 1 * time.Minute, nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookie.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
			},
			wantErr: nil,
			wantVars: func() step.Variables {
				vars := step.NewVariables()
				vars.Set(
					cookie.SessionName,
					"0123456789abcdef0123456789abcdef",
				)
				vars.Set(
					consts.SessionTTL,
					1*time.Minute,
				)
				vars.Set(
					consts.SessionCreatedAtField,
					timex.MustParse(time.RFC3339, "2026-01-01T00:00:00Z"),
				)
				return vars
			}(),
			wantPanic: false,
		},
		{
			name: "missing session in vars",
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
			name: "redis hgetall failed",
			step: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return nil, assert.AnError
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookie.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
			},
			wantErr: assert.AnError,
			wantVars: func() step.Variables {
				vars := step.NewVariables()
				vars.Set(
					cookie.SessionName,
					"0123456789abcdef0123456789abcdef",
				)
				return vars
			}(),
			wantPanic: false,
		},
		{
			name: "missed session creation time field",
			step: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{}, nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookie.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
			},
			wantErr: errs.ErrExpectationFailed,
			wantVars: func() step.Variables {
				vars := step.NewVariables()
				vars.Set(
					cookie.SessionName,
					"0123456789abcdef0123456789abcdef",
				)
				return vars
			}(),
			wantPanic: false,
		},
		{
			name: "session creation time is empty",
			step: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								"created_at": "",
							}, nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookie.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
			},
			wantErr: errs.ErrExpectationFailed,
			wantVars: func() step.Variables {
				vars := step.NewVariables()
				vars.Set(
					cookie.SessionName,
					"0123456789abcdef0123456789abcdef",
				)
				return vars
			}(),
			wantPanic: false,
		},
		{
			name: "session creation time is about a zero",
			step: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								"created_at": "0001-01-01T00:00:00Z",
							}, nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookie.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
			},
			wantErr: errs.ErrExpectationFailed,
			wantVars: func() step.Variables {
				vars := step.NewVariables()
				vars.Set(
					cookie.SessionName,
					"0123456789abcdef0123456789abcdef",
				)
				return vars
			}(),
			wantPanic: false,
		},
		{
			name: "session creation time is in the future",
			step: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								"created_at": "9999-12-31T23:59:59Z",
							}, nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookie.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
			},
			wantErr: errs.ErrExpectationFailed,
			wantVars: func() step.Variables {
				vars := step.NewVariables()
				vars.Set(
					cookie.SessionName,
					"0123456789abcdef0123456789abcdef",
				)
				return vars
			}(),
			wantPanic: false,
		},
		{
			name: "redis TTL failed",
			step: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								"created_at": "2026-01-01T00:00:00Z",
							}, nil
						},
					),
					redisfk.WithTTL(
						func(
							_ context.Context,
							_ string,
						) (time.Duration, error) {
							return 0, assert.AnError
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookie.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
			},
			wantErr: errs.ErrRedisFailed,
			wantVars: func() step.Variables {
				vars := step.NewVariables()
				vars.Set(
					cookie.SessionName,
					"0123456789abcdef0123456789abcdef",
				)
				return vars
			}(),
			wantPanic: false,
		},
		{
			name: "invalid session",
			step: redis.NewStep(
				redisfk.NewFakeClient(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookie.SessionName,
						"21321",
					)
					return vars
				}(),
			},
			wantErr: errs.ErrExpectationFailed,
			wantVars: func() step.Variables {
				vars := step.NewVariables()
				vars.Set(
					cookie.SessionName,
					"21321",
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
