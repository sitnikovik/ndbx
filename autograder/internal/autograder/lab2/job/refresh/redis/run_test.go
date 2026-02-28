package redis_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	cookieconsts "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/cookie"
	redisconsts "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/refresh/redis"
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
	type want struct {
		vars  step.Variables
		err   error
		panic bool
	}
	tests := []struct {
		name string
		step *redis.Step
		args args
		want want
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
							return time.Hour, nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					vars.Set(
						redisconsts.SessionCreatedAtField,
						timex.MustParse(
							time.RFC3339,
							"2026-01-01T00:00:00Z",
						),
					)
					vars.Set(
						redisconsts.SessionTTL,
						time.Hour,
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					vars.Set(
						redisconsts.SessionCreatedAtField,
						timex.MustParse(
							time.RFC3339,
							"2026-01-01T00:00:00Z",
						),
					)
					vars.Set(
						redisconsts.SessionTTL,
						time.Hour,
					)
					return vars
				}(),
				err:   nil,
				panic: false,
			},
		},
		{
			name: "invalid session id",
			step: redis.NewStep(
				redisfk.NewFakeClient(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookieconsts.SessionName,
						"invalid-session-id",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookieconsts.SessionName,
						"invalid-session-id",
					)
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "session not found in redis",
			step: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return nil, errors.New("session not found")
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
				err:   errs.ErrRedisFailed,
				panic: false,
			},
		},
		{
			name: "session created_at field is missing",
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
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "invalid time format for created_at field",
			step: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								"created_at": "invalid-time-format",
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
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "session created_at field is empty",
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
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "session creation date is zero",
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
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "created_at field in redis is updated",
			step: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								"created_at": "2026-01-01T01:00:00Z",
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
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					vars.Set(
						redisconsts.SessionCreatedAtField,
						timex.MustParse(
							time.RFC3339,
							"2026-01-01T00:00:00Z",
						),
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					vars.Set(
						redisconsts.SessionCreatedAtField,
						timex.MustParse(
							time.RFC3339,
							"2026-01-01T00:00:00Z",
						),
					)
					return vars
				}(),
				err: errs.ErrExpectationFailed,
			},
		},
		{
			name: "redis ttl failed",
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
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					vars.Set(
						redisconsts.SessionCreatedAtField,
						timex.MustParse(
							time.RFC3339,
							"2026-01-01T00:00:00Z",
						),
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					vars.Set(
						redisconsts.SessionCreatedAtField,
						timex.MustParse(
							time.RFC3339,
							"2026-01-01T00:00:00Z",
						),
					)
					return vars
				}(),
				err:   errs.ErrRedisFailed,
				panic: false,
			},
		},
		{
			name: "session ttl is not updated",
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
							return time.Minute, nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					vars.Set(
						redisconsts.SessionCreatedAtField,
						timex.MustParse(
							time.RFC3339,
							"2026-01-01T00:00:00Z",
						),
					)
					vars.Set(
						redisconsts.SessionTTL,
						time.Hour,
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						cookieconsts.SessionName,
						"0123456789abcdef0123456789abcdef",
					)
					vars.Set(
						redisconsts.SessionCreatedAtField,
						timex.MustParse(
							time.RFC3339,
							"2026-01-01T00:00:00Z",
						),
					)
					vars.Set(
						redisconsts.SessionTTL,
						time.Hour,
					)
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.want.panic {
				assert.Panics(t, func() {
					_ = tt.step.Run(
						tt.args.ctx,
						tt.args.vars,
					)
				})
				return
			}
			assert.ErrorIs(
				t,
				tt.step.Run(
					tt.args.ctx,
					tt.args.vars,
				),
				tt.want.err,
			)
			assert.Equal(
				t,
				tt.want.vars,
				tt.args.vars,
			)
		})
	}
}
