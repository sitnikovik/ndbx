package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/redis/user/recommendations/field"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/user/one/recommendations/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/step/user/one/recommendations/redis/expect"
	"github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/redis"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
)

func TestStep_Run(t *testing.T) {
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
		s    *impl.Step
		args args
		want want
	}{
		{
			name: "ok",
			s: impl.NewStep(
				NewDescFx(),
				redis.NewFakeClient(
					redis.WithHGetAll(
						func(_ context.Context, _ string) (map[string]string, error) {
							return map[string]string{
								field.Events: `[` +
									`{` +
									`"id": "1",` +
									`"title": "test title",` +
									`"description": "test description",` +
									`"location": {` +
									`"address": "test location"` +
									`},` +
									`"created_at": "2024-01-01T00:00:00Z",` +
									`"created_by": "test_user",` +
									`"started_at": "2024-01-01T01:00:00Z",` +
									`"finished_at": "2024-01-01T02:00:00Z"` +
									`}` +
									`]`,
							}, nil
						},
					),
					redis.WithTTL(
						func(_ context.Context, _ string) (time.Duration, error) {
							return 1 * time.Minute, nil
						},
					),
				),
				userfx.NewJohnDoe(),
				NewExpectationsFx(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
				err:   nil,
				panic: false,
			},
		},
		{
			name: "has no user hash in vars",
			s: impl.NewStep(
				NewDescFx(),
				redis.NewFakeClient(),
				userfx.NewJohnDoe(),
				NewExpectationsFx(),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   nil,
				panic: true,
			},
		},
		{
			name: "failed to hgetall",
			s: impl.NewStep(
				NewDescFx(),
				redis.NewFakeClient(
					redis.WithHGetAll(
						func(_ context.Context, _ string) (map[string]string, error) {
							return nil, assert.AnError
						},
					),
				),
				userfx.NewJohnDoe(),
				NewExpectationsFx(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
				err:   errs.ErrExternalDependencyFailed,
				panic: false,
			},
		},
		{
			name: "has no events field",
			s: impl.NewStep(
				NewDescFx(),
				redis.NewFakeClient(
					redis.WithHGetAll(
						func(_ context.Context, _ string) (map[string]string, error) {
							return map[string]string{}, nil
						},
					),
				),
				userfx.NewJohnDoe(),
				NewExpectationsFx(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "unexpected list",
			s: impl.NewStep(
				NewDescFx(),
				redis.NewFakeClient(
					redis.WithHGetAll(
						func(_ context.Context, _ string) (map[string]string, error) {
							return map[string]string{
								field.Events: `[]`,
							}, nil
						},
					),
				),
				userfx.NewJohnDoe(),
				NewExpectationsFx(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "wrong json in events",
			s: impl.NewStep(
				NewDescFx(),
				redis.NewFakeClient(
					redis.WithHGetAll(
						func(_ context.Context, _ string) (map[string]string, error) {
							return map[string]string{
								field.Events: `not json`,
							}, nil
						},
					),
				),
				userfx.NewJohnDoe(),
				NewExpectationsFx(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
				err:   errs.ErrMarshallFailed,
				panic: false,
			},
		},
		{
			name: "no events but expected",
			s: impl.NewStep(
				NewDescFx(),
				redis.NewFakeClient(
					redis.WithHGetAll(
						func(_ context.Context, _ string) (map[string]string, error) {
							return map[string]string{
								field.Events: `[]`,
							}, nil
						},
					),
					redis.WithTTL(
						func(_ context.Context, _ string) (time.Duration, error) {
							return 1 * time.Minute, nil
						},
					),
				),
				userfx.NewJohnDoe(),
				expect.NewExpectations(
					expect.WithNoEvents(),
					expect.WithTTL(1*time.Minute),
				),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
				err:   nil,
				panic: false,
			},
		},
		{
			name: "failed to get ttl",
			s: impl.NewStep(
				NewDescFx(),
				redis.NewFakeClient(
					redis.WithHGetAll(
						func(_ context.Context, _ string) (map[string]string, error) {
							return map[string]string{
								field.Events: `[` +
									`{` +
									`"id": "1",` +
									`"title": "test title",` +
									`"description": "test description",` +
									`"location": {` +
									`"address": "test location"` +
									`},` +
									`"created_at": "2024-01-01T00:00:00Z",` +
									`"created_by": "test_user",` +
									`"started_at": "2024-01-01T01:00:00Z",` +
									`"finished_at": "2024-01-01T02:00:00Z"` +
									`}` +
									`]`,
							}, nil
						},
					),
					redis.WithTTL(
						func(_ context.Context, _ string) (time.Duration, error) {
							return 0, assert.AnError
						},
					),
				),
				userfx.NewJohnDoe(),
				NewExpectationsFx(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
				err:   errs.ErrExternalDependencyFailed,
				panic: false,
			},
		},
		{
			name: "unexpected ttl",
			s: impl.NewStep(
				NewDescFx(),
				redis.NewFakeClient(
					redis.WithHGetAll(
						func(_ context.Context, _ string) (map[string]string, error) {
							return map[string]string{
								field.Events: `[` +
									`{` +
									`"id": "1",` +
									`"title": "test title",` +
									`"description": "test description",` +
									`"location": {` +
									`"address": "test location"` +
									`},` +
									`"created_at": "2024-01-01T00:00:00Z",` +
									`"created_by": "test_user",` +
									`"started_at": "2024-01-01T01:00:00Z",` +
									`"finished_at": "2024-01-01T02:00:00Z"` +
									`}` +
									`]`,
							}, nil
						},
					),
					redis.WithTTL(
						func(_ context.Context, _ string) (time.Duration, error) {
							return 2 * time.Minute, nil
						},
					),
				),
				userfx.NewJohnDoe(),
				NewExpectationsFx(),
			),
			args: args{
				ctx: context.Background(),
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
					)
					return vars
				}(),
			},
			want: want{
				vars: func() step.Variables {
					vars := step.NewVariables()
					vars.Set(
						userfx.NewJohnDoe().Hash(),
						"123",
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
					_ = tt.s.Run(
						tt.args.ctx,
						tt.args.vars,
					)
				})
				assert.Equal(
					t,
					tt.want.vars,
					tt.args.vars,
				)
				return
			}
			assert.ErrorIs(
				t,
				tt.s.Run(
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
