package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/cookie"
	consts "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/consts/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab2/job/preserve/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	redisfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

// NewValidStepVariables returns step variables for which step seems to be succeeded and not panic.
func NewValidStepVariables() step.Variables {
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
}

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
				ctx:  context.Background(),
				vars: NewValidStepVariables(),
			},
			wantErr:   nil,
			wantVars:  NewValidStepVariables(),
			wantPanic: false,
		},
		{
			name: "session missed in vars",
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
				ctx:  context.Background(),
				vars: NewValidStepVariables(),
			},
			wantErr:   errs.ErrRedisFailed,
			wantVars:  NewValidStepVariables(),
			wantPanic: false,
		},
		{
			name: "missed session creation time in response",
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
				ctx:  context.Background(),
				vars: NewValidStepVariables(),
			},
			wantErr:   errs.ErrExpectationFailed,
			wantVars:  NewValidStepVariables(),
			wantPanic: false,
		},
		{
			name: "invalid creation time",
			step: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								"created_at": "2026-01-01 00:00:00",
							}, nil
						},
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: NewValidStepVariables(),
			},
			wantErr:   errs.ErrExpectationFailed,
			wantVars:  NewValidStepVariables(),
			wantPanic: false,
		},
		{
			name: "session creation time doesnt equal to expected one",
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
				ctx:  context.Background(),
				vars: NewValidStepVariables(),
			},
			wantErr:   errs.ErrExpectationFailed,
			wantVars:  NewValidStepVariables(),
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
				ctx:  context.Background(),
				vars: NewValidStepVariables(),
			},
			wantErr:   errs.ErrRedisFailed,
			wantVars:  NewValidStepVariables(),
			wantPanic: false,
		},
		{
			name: "session TTL is less than expected",
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
							return 30 * time.Second, nil
						},
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: NewValidStepVariables(),
			},
			wantErr:   errs.ErrExpectationFailed,
			wantVars:  NewValidStepVariables(),
			wantPanic: false,
		},
		{
			name: "session TTL is greater than expected",
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
							return 2 * time.Minute, nil
						},
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: NewValidStepVariables(),
			},
			wantErr:   errs.ErrExpectationFailed,
			wantVars:  NewValidStepVariables(),
			wantPanic: false,
		},
		{
			name: "session TTL is in acceptable range",
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
							return 1*time.Minute + 3*time.Second, nil
						},
					),
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: NewValidStepVariables(),
			},
			wantErr:   nil,
			wantVars:  NewValidStepVariables(),
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
