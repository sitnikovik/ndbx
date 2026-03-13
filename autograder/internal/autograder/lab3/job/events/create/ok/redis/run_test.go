package redis_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	sessionCookies "github.com/sitnikovik/ndbx/autograder/internal/app/cookie/session"
	"github.com/sitnikovik/ndbx/autograder/internal/app/redis/user/session"
	"github.com/sitnikovik/ndbx/autograder/internal/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/lab3/job/events/create/ok/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/autograder/variable"
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
	type want struct {
		vars  step.Variables
		err   error
		panic bool
	}
	tests := []struct {
		name string
		s    *redis.Step
		args args
		want want
	}{
		{
			name: "ok",
			s: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								session.CreatedAtField: "2024-06-01T00:00:00Z",
								session.UpdatedAtField: "2024-06-01T00:15:00Z",
								session.UserIDField:    "1",
							}, nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(sessionCookies.Name, "test-session-id")
					vars.Set(variable.User, user.NewUser(
						"1",
						"sams3piol",
						"Sam Sepiol",
					))
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(sessionCookies.Name, "test-session-id")
					vars.Set(variable.User, user.NewUser(
						"1",
						"sams3piol",
						"Sam Sepiol",
					))
					return vars
				}(),
				err:   nil,
				panic: false,
			},
		},
		{
			name: "failed to get redis data",
			s: redis.NewStep(
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
					vars.Set(sessionCookies.Name, "test-session-id")
					vars.Set(variable.User, user.NewUser(
						"1",
						"sams3piol",
						"Sam Sepiol",
					))
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(sessionCookies.Name, "test-session-id")
					vars.Set(variable.User, user.NewUser(
						"1",
						"sams3piol",
						"Sam Sepiol",
					))
					return vars
				}(),
				err:   errs.ErrExternalDependencyFailed,
				panic: false,
			},
		},
		{
			name: "session updatedAt goes before createdAt",
			s: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								session.CreatedAtField: "2024-06-01T00:00:00Z",
								session.UpdatedAtField: "2024-05-31T23:59:59Z",
								session.UserIDField:    "1",
							}, nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(sessionCookies.Name, "test-session-id")
					vars.Set(variable.User, user.NewUser(
						"1",
						"sams3piol",
						"Sam Sepiol",
					))
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(sessionCookies.Name, "test-session-id")
					vars.Set(variable.User, user.NewUser(
						"1",
						"sams3piol",
						"Sam Sepiol",
					))
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "user id in session does not match user id in variables",
			s: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								session.CreatedAtField: "2024-06-01T00:00:00Z",
								session.UpdatedAtField: "2024-06-01T00:00:00Z",
								session.UserIDField:    "2",
							}, nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(sessionCookies.Name, "test-session-id")
					vars.Set(variable.User, user.NewUser(
						"1",
						"sams3piol",
						"Sam Sepiol",
					))
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(sessionCookies.Name, "test-session-id")
					vars.Set(variable.User, user.NewUser(
						"1",
						"sams3piol",
						"Sam Sepiol",
					))
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "session recreated",
			s: redis.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								session.CreatedAtField: "2024-06-01T00:00:00Z",
								session.UpdatedAtField: "2024-06-01T00:00:00Z",
								session.UserIDField:    "1",
							}, nil
						},
					),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(sessionCookies.Name, "test-session-id")
					vars.Set(variable.User, user.NewUser(
						"1",
						"sams3piol",
						"Sam Sepiol",
					))
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(sessionCookies.Name, "test-session-id")
					vars.Set(variable.User, user.NewUser(
						"1",
						"sams3piol",
						"Sam Sepiol",
					))
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
					_ = tt.s.Run(
						tt.args.ctx,
						tt.args.vars,
					)
				})
				return
			}
			assert.Equal(
				t,
				tt.args.vars,
				tt.want.vars,
			)
			assert.ErrorIs(
				t,
				tt.s.Run(
					tt.args.ctx,
					tt.args.vars,
				),
				tt.want.err,
			)
		})
	}
}
