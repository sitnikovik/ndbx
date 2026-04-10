package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/events/one/dislike/redis"
	redisfk "github.com/sitnikovik/ndbx/autograder/internal/test/fake/client/redis"
	eventfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/event"
	userfx "github.com/sitnikovik/ndbx/autograder/internal/test/fixture/app/user"
	"github.com/sitnikovik/ndbx/autograder/internal/timex"
)

var (
	// eventFixture is a fixture used in the tests cases.
	eventFixture = eventfx.NewBirthdayParty(
		event.NewDates(
			timex.MustRFC3339("2026-03-31T15:00:00Z"),
			timex.MustRFC3339("2026-03-31T23:00:00Z"),
		),
		timex.MustRFC3339("2026-03-14T12:31:00Z"),
		userfx.NewSamwiseGamgee(),
	)
	// vars is the step variables used in the tests cases.
	vars = func() step.Variables {
		vv := step.NewVariables()
		vv.Set(
			eventFixture.Hash(),
			"13298",
		)
		return vv
	}()
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
		s    *impl.Step
		args args
		want want
	}{
		{
			name: "ok",
			s: impl.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGet(
						func(
							_ context.Context,
							_, _ string,
						) (string, error) {
							return "3", nil
						},
					),
					redisfk.WithTTL(
						func(
							_ context.Context,
							_ string,
						) (time.Duration, error) {
							return 2 * time.Second, nil
						},
					),
				),
				eventFixture,
				3,
				2*time.Second,
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
				err:   nil,
				panic: false,
			},
		},
		{
			name: "empty vars",
			s: impl.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGet(
						func(
							_ context.Context,
							_, _ string,
						) (string, error) {
							return "3", nil
						},
					),
					redisfk.WithTTL(
						func(
							_ context.Context,
							_ string,
						) (time.Duration, error) {
							return 2 * time.Second, nil
						},
					),
				),
				eventFixture,
				3,
				2*time.Second,
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   nil,
				panic: false,
			},
		},
		{
			name: "failed to get reactions",
			s: impl.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGet(
						func(
							_ context.Context,
							_, _ string,
						) (string, error) {
							return "", assert.AnError
						},
					),
				),
				eventFixture,
				3,
				2*time.Second,
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
				err:   errs.ErrExternalDependencyFailed,
				panic: false,
			},
		},
		{
			name: "got empty string in likes field",
			s: impl.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGet(
						func(
							_ context.Context,
							_, _ string,
						) (string, error) {
							return "", nil
						},
					),
				),
				eventFixture,
				3,
				2*time.Second,
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "likes do not meets expectations",
			s: impl.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGet(
						func(
							_ context.Context,
							_, _ string,
						) (string, error) {
							return "1", nil
						},
					),
				),
				eventFixture,
				3,
				2*time.Second,
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "failed to get ttl",
			s: impl.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGet(
						func(
							_ context.Context,
							_, _ string,
						) (string, error) {
							return "3", nil
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
				eventFixture,
				3,
				2*time.Second,
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
				err:   errs.ErrExternalDependencyFailed,
				panic: false,
			},
		},
		{
			name: "got unexpected ttl",
			s: impl.NewStep(
				redisfk.NewFakeClient(
					redisfk.WithHGet(
						func(
							_ context.Context,
							_, _ string,
						) (string, error) {
							return "3", nil
						},
					),
					redisfk.WithTTL(
						func(
							_ context.Context,
							_ string,
						) (time.Duration, error) {
							return 3 * time.Second, nil
						},
					),
				),
				eventFixture,
				3,
				2*time.Second,
			),
			args: args{
				ctx:  context.Background(),
				vars: vars,
			},
			want: want{
				vars:  vars,
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
