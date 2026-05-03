package redis_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/ndbx/autograder/internal/app/event"
	"github.com/sitnikovik/ndbx/autograder/internal/app/rating"
	"github.com/sitnikovik/ndbx/autograder/internal/app/redis/event/reviews/field"
	"github.com/sitnikovik/ndbx/autograder/internal/app/review/count"
	"github.com/sitnikovik/ndbx/autograder/internal/errs"
	"github.com/sitnikovik/ndbx/autograder/internal/step"
	impl "github.com/sitnikovik/ndbx/autograder/internal/step/events/one/review/redis"
	"github.com/sitnikovik/ndbx/autograder/internal/step/events/one/review/redis/expect"
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
	expectFx = expect.NewExpectations(
		expect.WithCounts(
			count.NewCounts(
				count.WithRating(
					rating.NewRating(4.8),
				),
				count.WithCount(3),
			),
		),
		expect.WithTTL(2*time.Second),
	)
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
				step.NewDesc(
					"test title",
					"test description",
				),
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								field.Rating: "4.8",
								field.Count:  "3",
							}, nil
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
				impl.WithExpectations(
					expectFx,
				),
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
			name: "failed to hget all",
			s: impl.NewStep(
				step.NewDesc(
					"test title",
					"test description",
				),
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
				eventFixture,
				impl.WithExpectations(
					expectFx,
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrExternalDependencyFailed,
				panic: false,
			},
		},
		{
			name: "got no rating from redis",
			s: impl.NewStep(
				step.NewDesc(
					"test title",
					"test description",
				),
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								field.Count: "3",
							}, nil
						},
					),
				),
				eventFixture,
				impl.WithExpectations(
					expectFx,
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "got unexpected rating",
			s: impl.NewStep(
				step.NewDesc(
					"test title",
					"test description",
				),
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								field.Rating: "4.9",
							}, nil
						},
					),
				),
				eventFixture,
				impl.WithExpectations(
					expectFx,
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "got no count from redis",
			s: impl.NewStep(
				step.NewDesc(
					"test title",
					"test description",
				),
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								field.Rating: "4.8",
							}, nil
						},
					),
				),
				eventFixture,
				impl.WithExpectations(
					expectFx,
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "got unexpected count",
			s: impl.NewStep(
				step.NewDesc(
					"test title",
					"test description",
				),
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								field.Rating: "4.8",
								field.Count:  "123",
							}, nil
						},
					),
				),
				eventFixture,
				impl.WithExpectations(
					expectFx,
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrExpectationFailed,
				panic: false,
			},
		},
		{
			name: "failed to get ttl",
			s: impl.NewStep(
				step.NewDesc(
					"test title",
					"test description",
				),
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								field.Rating: "4.8",
								field.Count:  "3",
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
				eventFixture,
				impl.WithExpectations(
					expectFx,
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
				err:   errs.ErrExternalDependencyFailed,
				panic: false,
			},
		},
		{
			name: "got unexpected ttl",
			s: impl.NewStep(
				step.NewDesc(
					"test title",
					"test description",
				),
				redisfk.NewFakeClient(
					redisfk.WithHGetAll(
						func(
							_ context.Context,
							_ string,
						) (map[string]string, error) {
							return map[string]string{
								field.Rating: "4.8",
								field.Count:  "3",
							}, nil
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
				impl.WithExpectations(
					expectFx,
				),
			),
			args: args{
				ctx:  context.Background(),
				vars: step.NewVariables(),
			},
			want: want{
				vars:  step.NewVariables(),
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
